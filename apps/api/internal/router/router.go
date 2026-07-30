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
	mux.HandleFunc("/api/v1/metrics", handlers.MetricsHandler)
	mux.HandleFunc("/api/v1/supply", handlers.SupplyHandler)
	mux.HandleFunc("/api/v1/hashrate", handlers.HashrateHandler)

	return middleware.Logging(mux)
}
