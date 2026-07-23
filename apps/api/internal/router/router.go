package router

import (
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/handlers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)

	return mux
}