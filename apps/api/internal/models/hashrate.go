package models

type HashrateInfo struct {
	Hashrate float64 `json:"hashrate"`

	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}
