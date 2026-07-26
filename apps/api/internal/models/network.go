package models

type NetworkInfo struct {
	Network         string `json:"network"`
	BlockCount      uint64 `json:"block_count"`
	HeaderCount     uint64 `json:"header_count"`
	Difficulty      uint64 `json:"difficulty"`
	VirtualDAAScore uint64 `json:"virtual_daa_score"`

	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}
