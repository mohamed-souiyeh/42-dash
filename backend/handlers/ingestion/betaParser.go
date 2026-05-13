package ingestion

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// ProviderBStruct represents the raw JSON from Provider 2
type ProviderBStruct struct {
	GameCode     string `json:"gameCode"`
	GameName     string `json:"gameName"`
	ProviderName string `json:"provideName"`
	GameCategory string `json:"gameCategory"` // e.g., "SL"
	RTPValue     string `json:"rtpValue"`     // e.g., "95.42"
	RiskLevel    string `json:"riskLevel"`    // e.g., "MED"
	IsEnabled    int    `json:"isEnabled"`    // 1 or 0
	ReleaseDate  string `json:"releaseDate"`  // e.g., "15/03/2022"
	ImageURL     string `json:"imageUrl"`
	TagList      string `json:"tagList"` // e.g., "megaways,free-spins"
}

type BetaAdapter struct{}

func (a *BetaAdapter) Parse(rawJSON []byte) ([]GamePayload, error) {
	var rawGames []ProviderBStruct
	if err := json.Unmarshal(rawJSON, &rawGames); err != nil {
		return nil, err
	}

	// log.Info("the raw games: ", "rawgames", rawGames)
	// log.Info("the raw body: ", "rawbody", string(rawJSON))

	varianceMap := map[string]string{
		"LOW":  "Low",
		"MED":  "Mid",
		"HIGH": "High",
	}

	categoryMap := map[string]string{
		"SL": "Slots",
		"LV": "Live",
		"TB": "Table",
		"IN": "Instant",
		"JP": "Jackpot",
	}

	var cleanGames []GamePayload
	for _, raw := range rawGames {
		layout := "02/01/2006"
		// log.Info("the raw date: ", "rawdate", raw.ReleaseDate)
		// log.Info("the raw data: ", "rawdata", raw)
		parsedTime, err := time.Parse(layout, raw.ReleaseDate)
		if err != nil {
			return nil, err
		}

		parsedRTP, err := strconv.ParseFloat(raw.RTPValue, 64)
		if err != nil {
			return nil, err
		}

		// Map the weird data to your clean standard format
		clean := GamePayload{
			ID:         raw.GameCode,
			Name:       raw.GameName,
			Provider:   raw.ProviderName,
			Category:   categoryMap[raw.GameCategory],
			RTP:        parsedRTP,
			Variance:   varianceMap[raw.RiskLevel],
			Enabled:    raw.IsEnabled != 0,
			LaunchDate: parsedTime,
			ImageURL:   raw.ImageURL,
			Tags:       strings.Split(raw.TagList, ","),
		}
		cleanGames = append(cleanGames, clean)
	}

	return cleanGames, nil
}
