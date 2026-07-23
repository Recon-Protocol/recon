package services

type NetworkResponse struct {
	Network   string `json:"network"`
	APIVersion string `json:"api_version"`
	Service   string `json:"service"`
	Status    string `json:"status"`
}

func GetNetwork() NetworkResponse {
	return NetworkResponse{
		Network:    "kaspa-mainnet",
		APIVersion: "v1",
		Service:    "recon-api",
		Status:     "online",
	}
}