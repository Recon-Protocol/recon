package kaspa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Recon-Protocol/recon/apps/api/internal/models"
)

type networkAPIResponse struct {
	NetworkName     string      `json:"networkName"`
	BlockCount      interface{} `json:"blockCount"`
	HeaderCount     interface{} `json:"headerCount"`
	Difficulty      interface{} `json:"difficulty"`
	VirtualDAAScore interface{} `json:"virtualDaaScore"`
}

func (c *Client) GetNetworkInfo() (models.NetworkInfo, error) {

	req, err := http.NewRequest(http.MethodGet, networkEndpoint, nil)
	if err != nil {
		return models.NetworkInfo{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "RECON/0.1")

	resp, err := c.http.Do(req)
	if err != nil {
		return models.NetworkInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.NetworkInfo{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var api networkAPIResponse

	if err := json.NewDecoder(resp.Body).Decode(&api); err != nil {
		return models.NetworkInfo{}, err
	}

	return models.NetworkInfo{
		Network:         api.NetworkName,
		BlockCount:      toUint64(api.BlockCount),
		HeaderCount:     toUint64(api.HeaderCount),
		Difficulty:      toUint64(api.Difficulty),
		VirtualDAAScore: toUint64(api.VirtualDAAScore),
		Status:          "online",
		Service:         "recon-api",
		APIVersion:      "v1",
	}, nil
}

func toUint64(v interface{}) uint64 {

	switch value := v.(type) {

	case float64:
		return uint64(value)

	case string:
		n, _ := strconv.ParseUint(value, 10, 64)
		return n

	case json.Number:
		n, _ := value.Int64()
		return uint64(n)
	}

	return 0
}
