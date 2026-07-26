package kaspa

import (
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/client"
)

const networkEndpoint = "https://api.kaspa.org/info/network"

type Client struct {
	http *http.Client
}

func New() *Client {
	return &Client{
		http: client.NewHTTPClient(),
	}
}
