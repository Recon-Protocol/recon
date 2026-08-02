package services

import (
	"fmt"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetReport() (models.Report, error) {

	analytics, err := GetAnalytics()
	if err != nil {
		return models.Report{}, err
	}

	network, err := GetNetwork()
	if err != nil {
		return models.Report{}, err
	}

	report := models.Report{
		Network:           network.Network,
		Hashrate:          fmt.Sprintf("%.2f PH/s", analytics.HashratePHS),
		CirculatingSupply: fmt.Sprintf("%.2fB KAS", float64(analytics.CirculatingSupply)/100000000/1000000000),
		PercentMined:      fmt.Sprintf("%.2f%%", analytics.PercentMined),
		BlockReward:       analytics.BlockReward + " KAS",
		NextReduction:     fmt.Sprintf("%.1f days", analytics.DaysToReduction),
		NetworkStatus:     "Healthy",

		Status:     analytics.Status,
		Service:    analytics.Service,
		APIVersion: analytics.APIVersion,
	}

	return report, nil
}
