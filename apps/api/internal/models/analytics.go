package models

type Analytics struct {
	HashratePHS float64 `json:"hashrate_phs"`

	CirculatingSupply uint64 `json:"circulating_supply"`
	MaxSupply         uint64 `json:"max_supply"`

	PercentMined    float64 `json:"percent_mined"`
	RemainingSupply uint64  `json:"remaining_supply"`

	BlockReward     string  `json:"block_reward"`
	NextReduction   string  `json:"next_reduction"`
	DaysToReduction float64 `json:"days_to_reduction"`

	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}
