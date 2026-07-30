package services

import (
	"github.com/Recon-Protocol/recon/apps/api/internal/client/kaspa"
	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetSupply() (models.SupplyInfo, error) {

	client := kaspa.New()

	supply, err := client.GetSupplyInfo()
	if err != nil {
		return models.SupplyInfo{}, err
	}

	return supply, nil
}
