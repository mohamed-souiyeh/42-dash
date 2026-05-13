package ingestion

import (
	"encoding/json"
	"time"
)

type ProviderAStruct struct {
	GameID         string   `json:"gameId"`
	Title          string   `json:"title"`
	Studio         string   `json:"studio"`
	Type           string   `json:"type"`
	ReturnToPlayer float64  `json:"returnToPlayer"`
	Variance       string   `json:"variance"` // "medium"
	Active         bool     `json:"active"`
	LaunchDate     string   `json:"launchDate"` // "2022-03-15T00:00:00Z"
	Thumbnail      string   `json:"thumbnail"`
	Features       []string `json:"features"`
}

type AlphaAdapter struct{}

func (a *AlphaAdapter) Parse(rawJSON []byte) ([]GamePayload, error) {
	var rawGames []ProviderAStruct
	if err := json.Unmarshal(rawJSON, &rawGames); err != nil {
		return nil, err
	}

	varianceMap := map[string]string{
		"low":    "Low",
		"medium": "Mid",
		"high":   "High",
	}

	categoryMap := map[string]string{
		"slots":   "Slots",
		"live":    "Live",
		"table":   "Table",
		"instant": "Instant",
		"jackpot": "Jackpot",
	}

	var cleanGames []GamePayload
	for _, raw := range rawGames {
		parsedTime, err := time.Parse(time.RFC3339, raw.LaunchDate)
		if err != nil {
			return nil, err
		}
		// Map the weird data to your clean standard format
		clean := GamePayload{
			ID:         raw.GameID,
			Name:       raw.Title,
			Provider:   raw.Studio,
			Category:   categoryMap[raw.Type],
			RTP:        raw.ReturnToPlayer,
			Variance:   varianceMap[raw.Variance],
			Enabled:    raw.Active,
			LaunchDate: parsedTime,
			ImageURL:   raw.Thumbnail,
			Tags:       raw.Features,
		}
		cleanGames = append(cleanGames, clean)
	}

	return cleanGames, nil
}
