package handlers

import (
	repo "backend/db/repository"
	sqlc "backend/db/sqlc_generated"
	"backend/handlers/ingestion"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type IngestRequest struct {
	SourceURL string `json:"source_url"`
}

type IngestHandler struct {
	IngestRepo repo.IngestionRepository
	GameRepo   repo.GameRepository
}

func getAdapter(sourceURL string) (ingestion.ProviderAdapter, error) {
	switch {
	case strings.Contains(sourceURL, "alpha"):
		return &ingestion.AlphaAdapter{}, nil
	case strings.Contains(sourceURL, "beta"):
		return &ingestion.BetaAdapter{}, nil
	case strings.Contains(sourceURL, "gamma"):
		return &ingestion.GammaAdapter{}, nil
	default:
		return nil, errors.New("unknown provider")
	}
}

func (env *IngestHandler) HandleTriggerIngest(w http.ResponseWriter, r *http.Request) {
	var req IngestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	logEntry, err := env.IngestRepo.StartIngestion(ctx, req.SourceURL)
	if err != nil {
		http.Error(w, "Failed to initialize ingestion log", http.StatusInternalServerError)
		return
	}

	log.Info("log entry is: ", "entry", logEntry)

	if logEntry.Status == "SUCCESS" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data from this URL has already been successfully ingested."))
		return
	}

	// 3. Make the HTTP GET request to the mock server
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(req.SourceURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		// Mark as FAILED if the network request fails
		env.IngestRepo.UpdateIngestionStatus(ctx, sqlc.UpdateIngestionStatusParams{
			SourceUrl: req.SourceURL,
			Status:    "FAILED",
		})
		http.Error(w, "Failed to fetch data from source", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read the rawBody", http.StatusInternalServerError)
		return
	}

	adapter, err := getAdapter(req.SourceURL)
	if err != nil {
		http.Error(w, "unknown provider", http.StatusNotImplemented)
		return
	}

	cleanGames, err := adapter.Parse(rawBody)
	if err != nil {
		http.Error(w, "Failed to parse the raw Body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// we need to upsert the resulting cleanGames slice here
	for _, game := range cleanGames {
		gameTags, err := json.Marshal(game.Tags)
		if err != nil {
			http.Error(w, "failed to convert the tags to json", http.StatusInternalServerError)
			return
		}

		dbGame := sqlc.UpsertGameParams{
			ID:         game.ID,
			Name:       game.Name,
			Provider:   game.Provider,
			Category:   game.Category,
			Rtp:        game.RTP,
			Variance:   game.Variance,
			Enabled:    game.Enabled,
			LaunchDate: game.LaunchDate.Format(time.RFC3339),
			Tags:       string(gameTags),
			ImageUrl:   game.ImageURL,
		}
		_, err = env.GameRepo.UpsertGame(ctx, dbGame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 6. Success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ingestion completed successfully!"))
}
