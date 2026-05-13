package ingestion

import (
	"encoding/json"
	"time"
)

// ProviderCStruct represents the raw nested JSON from Provider 3
type ProviderCStruct struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			DisplayName string `json:"displayName"`
			Provider    struct {
				Code  string `json:"code"`
				Label string `json:"label"`
			} `json:"provider"`
			Classification struct {
				Category   string `json:"category"`
				Volatility string `json:"volatility"`
			} `json:"classification"`
			Metrics struct {
				RTP float64 `json:"rtp"` // e.g., 0.9542
			} `json:"metrics"`
			Status struct {
				Enabled  bool   `json:"enabled"`
				Released string `json:"released"` // e.g., "2022-03-15"
			} `json:"status"`
			Media struct {
				ThumbnailURL string `json:"thumbnailUrl"`
			} `json:"media"`
			Tags []struct {
				Slug string `json:"slug"`
			} `json:"tags"`
		} `json:"attributes"`
	} `json:"data"`
}

type GammaAdapter struct{}

func (a *GammaAdapter) Parse(rawJSON []byte) ([]GamePayload, error) {
	var rawGames []ProviderCStruct
	if err := json.Unmarshal(rawJSON, &rawGames); err != nil {
		return nil, err
	}

	// log.Info("the raw games: ", "rawgames", rawGames)
	// log.Info("the raw body: ", "rawbody", string(rawJSON))

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
		layout := "2006-01-02"
		// log.Info("the raw date: ", "rawdate", raw.ReleaseDate)
		// log.Info("the raw data: ", "rawdata", raw)
		parsedTime, err := time.Parse(layout, raw.Data.Attributes.Status.Released)
		if err != nil {
			return nil, err
		}

		parsedRTP := raw.Data.Attributes.Metrics.RTP * 100.00

		cleanTags := make([]string, 0, len(raw.Data.Attributes.Tags))

		for _, tagobj := range raw.Data.Attributes.Tags {
			cleanTags = append(cleanTags, tagobj.Slug)
		}

		// Map the weird data to your clean standard format
		clean := GamePayload{
			ID:         raw.Data.ID,
			Name:       raw.Data.Attributes.DisplayName,
			Provider:   raw.Data.Attributes.Provider.Label,
			Category:   categoryMap[raw.Data.Attributes.Classification.Category],
			RTP:        parsedRTP,
			Variance:   varianceMap[raw.Data.Attributes.Classification.Volatility],
			Enabled:    raw.Data.Attributes.Status.Enabled,
			LaunchDate: parsedTime,
			ImageURL:   raw.Data.Attributes.Media.ThumbnailURL,
			Tags:       cleanTags,
		}
		cleanGames = append(cleanGames, clean)
	}

	return cleanGames, nil
}
