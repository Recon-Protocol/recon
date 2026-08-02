package models

type Report struct {
	Network           string `json:"network"`
	Hashrate          string `json:"hashrate"`
	CirculatingSupply string `json:"circulating_supply"`
	PercentMined      string `json:"percent_mined"`
	BlockReward       string `json:"block_reward"`
	NextReduction     string `json:"next_reduction"`
	NetworkStatus     string `json:"network_status"`

	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}
