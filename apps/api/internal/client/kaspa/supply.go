package kaspa

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

type coinSupplyResponse struct {
	Circulating interface{} `json:"circulatingSupply"`
	MaxSupply   interface{} `json:"maxSupply"`
}

type blockRewardResponse struct {
	BlockReward interface{} `json:"blockreward"`
}

type halvingResponse struct {
	NextHalvingTimestamp uint64      `json:"nextHalvingTimestamp"`
	NextHalvingDate      string      `json:"nextHalvingDate"`
	NextHalvingAmount    interface{} `json:"nextHalvingAmount"`
}

func (c *Client) GetSupplyInfo() (models.SupplyInfo, error) {

	var supply models.SupplyInfo

	// ----------------------------------
	// Coin Supply
	// ----------------------------------

	req, err := http.NewRequest(
		http.MethodGet,
		baseURL+"/info/coinsupply",
		nil,
	)
	if err != nil {
		return supply, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "RECON/0.1")

	resp, err := c.http.Do(req)
	if err != nil {
		return supply, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return supply, fmt.Errorf("coinsupply: %s", resp.Status)
	}

	var coinResp coinSupplyResponse

	if err := json.NewDecoder(resp.Body).Decode(&coinResp); err != nil {
		return supply, err
	}

	supply.CirculatingSupply = toUint64(coinResp.Circulating)
	supply.TotalSupply = toUint64(coinResp.MaxSupply)

	// ----------------------------------
	// Block Reward
	// ----------------------------------

	req, err = http.NewRequest(
		http.MethodGet,
		baseURL+"/info/blockreward",
		nil,
	)
	if err != nil {
		return supply, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "RECON/0.1")

	resp, err = c.http.Do(req)
	if err != nil {
		return supply, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return supply, fmt.Errorf("blockreward: %s", resp.Status)
	}

	var rewardResp blockRewardResponse

	if err := json.NewDecoder(resp.Body).Decode(&rewardResp); err != nil {
		return supply, err
	}

	supply.BlockReward = fmt.Sprintf("%v", rewardResp.BlockReward)

	// ----------------------------------
	// Halving
	// ----------------------------------

	req, err = http.NewRequest(
		http.MethodGet,
		baseURL+"/info/halving",
		nil,
	)
	if err != nil {
		return supply, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "RECON/0.1")

	resp, err = c.http.Do(req)
	if err != nil {
		return supply, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return supply, fmt.Errorf("halving: %s", resp.Status)
	}

	var halvingResp halvingResponse

	if err := json.NewDecoder(resp.Body).Decode(&halvingResp); err != nil {
		return supply, err
	}

	supply.NextReduction = halvingResp.NextHalvingDate

	supply.Status = "online"
	supply.Service = "recon-api"
	supply.APIVersion = "v1"

	return supply, nil
}
