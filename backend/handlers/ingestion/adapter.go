package ingestion

import "time"

type GamePayload struct {
	ID         string    `json:"ID"`
	Name       string    `json:"name"`
	Provider   string    `json:"provider"`
	Category   string    `json:"category"`
	RTP        float64   `json:"RTP"`
	Variance   string    `json:"variance"`
	Enabled    bool      `json:"enabled"`
	LaunchDate time.Time `json:"launchDate"`
	Tags       []string  `json:"Tags"`
	ImageURL   string    `json:"imageURL"`
}

type ProviderAdapter interface {
	Parse(rawJSON []byte) ([]GamePayload, error)
}
