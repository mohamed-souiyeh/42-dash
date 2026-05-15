package handlers

import (
	repo "backend/db/repository"
	"database/sql"
	"encoding/json"
	"net/http"
)

type GameHandlers struct {
	GameRepository repo.GameRepository
}

func (env *GameHandlers) GetGame(w http.ResponseWriter, r *http.Request) {
	gameID := r.PathValue("id")
	if gameID == "" {
		http.Error(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	game, err := env.GameRepository.GetGameByID(ctx, gameID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return a 404 if the game doesn't exist
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}
		// Return a 500 for any other database errors
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(game); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
