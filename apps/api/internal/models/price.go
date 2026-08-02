package models

type PriceInfo struct {
	PriceUSD     float64 `json:"price_usd"`
	MarketCapUSD float64 `json:"market_cap_usd"`
	Volume24hUSD float64 `json:"volume_24h_usd"`
	Change24h    float64 `json:"change_24h"`

	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}
