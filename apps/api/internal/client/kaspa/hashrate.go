package kaspa

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

type hashrateResponse struct {
	Hashrate float64 `json:"hashrate"`
}

func (c *Client) GetHashrate() (models.HashrateInfo, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/info/hashrate",
		nil,
	)
	if err != nil {
		return models.HashrateInfo{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "RECON/0.1")

	resp, err := c.http.Do(req)
	if err != nil {
		return models.HashrateInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.HashrateInfo{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var api hashrateResponse

	if err := json.NewDecoder(resp.Body).Decode(&api); err != nil {
		return models.HashrateInfo{}, err
	}

	return models.HashrateInfo{
		Hashrate:   api.Hashrate,
		Status:     "online",
		Service:    "recon-api",
		APIVersion: "v1",
	}, nil
}
