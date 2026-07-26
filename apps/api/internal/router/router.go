package router

import (
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/handlers"
	"github.com/Recon-Protocol/recon/apps/api/internal/middleware"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// API v1
	mux.HandleFunc("/api/v1/health", handlers.HealthHandler)
	mux.HandleFunc("/api/v1/network", handlers.NetworkHandler)

	return middleware.Logging(mux)
}
