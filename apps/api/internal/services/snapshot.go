package services

import (
	"time"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

func GetSnapshot() (models.NetworkSnapshot, error) {

	network, err := GetNetwork()
	if err != nil {
		return models.NetworkSnapshot{}, err
	}

	return models.NetworkSnapshot{
		NetworkInfo: network,
		Timestamp:   time.Now().Unix(),
	}, nil
}
