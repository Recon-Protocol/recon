package models

type ExchangeFlow struct {
	Name      string  `json:"name"`
	BalanceKas float64 `json:"balance_kas"`
	Change24h  float64 `json:"change_24h"`
	Change7d   float64 `json:"change_7d"`
	Change30d  float64 `json:"change_30d"`
}

type Report struct {
	// Header
	ReportNumber string `json:"report_number"`
	Date         string `json:"date"`
	CalendarWeek string `json:"calendar_week"`
	Network      string `json:"network"`

	// Network (api.kaspa.org)
	Hashrate      string `json:"hashrate"`
	DAAScore      string `json:"daa_score"`
	NetworkStatus string `json:"network_status"`

	// Mining (api.kaspa.org)
	BlockReward   string `json:"block_reward"`
	NextReduction string `json:"next_reduction"`

	// Supply (api.kaspa.org)
	CirculatingSupply string `json:"circulating_supply"`
	PercentMined      string `json:"percent_mined"`

	// Exchange Flows (kaspa-lens.com)
	ExchangeFlows []ExchangeFlow `json:"exchange_flows"`

	// Manual fields (kaspa.stream)
	TPS        string `json:"tps"`
	BPS        string `json:"bps"`
	Nodes      string `json:"nodes"`
	MinerCount string `json:"miner_count"`
	PoolRevenue string `json:"pool_revenue"`
	L2Activity  string `json:"l2_activity"`

	// Metadata
	Status     string `json:"status"`
	Service    string `json:"service"`
	APIVersion string `json:"api_version"`
}