package kaspa

import (
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/client"
)

const baseURL = "https://api.kaspa.org"

type Client struct {
	http *http.Client
}

func New() *Client {
	return &Client{
		http: client.NewHTTPClient(),
	}
}
