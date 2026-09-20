package router

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Recon-Protocol/recon/apps/api/docs"
	"github.com/Recon-Protocol/recon/apps/api/internal/handlers"
	"github.com/Recon-Protocol/recon/apps/api/internal/middleware"
)

func New() http.Handler {
	mux := http.NewServeMux()

	// Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// API v1
	mux.HandleFunc("/api/v1/health", handlers.HealthHandler)
	mux.HandleFunc("/api/v1/network", handlers.NetworkHandler)
	mux.HandleFunc("/api/v1/metrics", handlers.MetricsHandler)
	mux.HandleFunc("/api/v1/supply", handlers.SupplyHandler)
	mux.HandleFunc("/api/v1/hashrate", handlers.HashrateHandler)
	mux.HandleFunc("/api/v1/price", handlers.PriceHandler)
	mux.HandleFunc("/api/v1/analytics", handlers.AnalyticsHandler)
	mux.HandleFunc("/api/v1/report", handlers.ReportHandler)
	mux.HandleFunc("/api/v1/intelligence", handlers.IntelligenceHandler)

	return middleware.CORS(middleware.Logging(mux))
}
