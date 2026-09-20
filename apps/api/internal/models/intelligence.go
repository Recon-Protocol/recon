package models

type IntelligenceComponent struct {
	Score int `json:"score"`
	Max   int `json:"max"`
}

type Intelligence struct {
	Score int    `json:"score"`
	Grade string `json:"grade"`
	Trend string `json:"trend"`

	Components struct {
		NetworkSecurity  IntelligenceComponent `json:"network_security"`
		SupplyState      IntelligenceComponent `json:"supply_state"`
		EmissionPressure IntelligenceComponent `json:"emission_pressure"`
		ProtocolEvent    IntelligenceComponent `json:"protocol_event"`
	} `json:"components"`

	Signals    []string `json:"signals"`
	Status     string   `json:"status"`
	Service    string   `json:"service"`
	APIVersion string   `json:"api_version"`
}
