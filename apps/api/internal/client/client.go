package client

import "github.com/Recon-Protocol/recon/apps/api/internal/models"

type NetworkClient interface {
	GetNetworkInfo() (models.NetworkInfo, error)
}
