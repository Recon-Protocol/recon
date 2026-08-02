package services

import (
	"time"

	"github.com/Recon-Protocol/recon/apps/api/internal/formatter"
	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetAnalytics() (models.Analytics, error) {

	network, err := GetNetwork()
	if err != nil {
		return models.Analytics{}, err
	}

	supply, err := GetSupply()
	if err != nil {
		return models.Analytics{}, err
	}

	hashrate, err := GetHashrate()
	if err != nil {
		return models.Analytics{}, err
	}

	nextReduction, err := time.Parse(
		"2006-01-02 15:04:05 MST",
		supply.NextReduction,
	)
	if err != nil {
		return models.Analytics{}, err
	}

	return models.Analytics{

		HashratePHS: ToPHS(hashrate.Hashrate),

		CirculatingSupply: supply.CirculatingSupply,
		MaxSupply:         supply.TotalSupply,

		PercentMined: PercentMined(
			supply.CirculatingSupply,
			supply.TotalSupply,
		),

		RemainingSupply: RemainingSupply(
			supply.CirculatingSupply,
			supply.TotalSupply,
		),

		BlockReward:   supply.BlockReward,
		NextReduction: supply.NextReduction,
		DaysToReduction: formatter.DaysBetween(
			time.Now(),
			nextReduction,
		),

		Status:     network.Status,
		Service:    network.Service,
		APIVersion: network.APIVersion,
	}, nil
}
