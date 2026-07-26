package services

import (
	"github.com/Recon-Protocol/recon/apps/api/internal/client/kaspa"
	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetNetwork() (models.NetworkInfo, error) {

	client := kaspa.New()

	return client.GetNetworkInfo()
}
