package services

import (
	"github.com/Recon-Protocol/recon/apps/api/internal/client/kaspa"
	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetHashrate() (models.HashrateInfo, error) {

	client := kaspa.New()

	return client.GetHashrate()
}
