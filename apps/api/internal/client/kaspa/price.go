package kaspa

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

type priceResponse struct {
	Price float64 `json:"price"`
}

func (c *Client) GetPriceInfo() (models.PriceInfo, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/info/price",
		nil,
	)
	if err != nil {
		return models.PriceInfo{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "RECON/0.1")

	resp, err := c.http.Do(req)
	if err != nil {
		return models.PriceInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.PriceInfo{}, fmt.Errorf("price: %s", resp.Status)
	}

	var api priceResponse

	if err := json.NewDecoder(resp.Body).Decode(&api); err != nil {
		return models.PriceInfo{}, err
	}

	return models.PriceInfo{
		PriceUSD:   api.Price,
		Status:     "online",
		Service:    "recon-api",
		APIVersion: "v1",
	}, nil
}
