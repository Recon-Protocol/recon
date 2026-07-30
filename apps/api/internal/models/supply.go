package models

type SupplyInfo struct {
	CirculatingSupply uint64 `json:"circulating_supply"`
	TotalSupply       uint64 `json:"total_supply"`
	BlockReward       string `json:"block_reward"`
	NextReduction     string `json:"next_reduction"`

	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}
