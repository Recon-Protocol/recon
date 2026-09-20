package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// AnalyticsHandler godoc
//
//	@Summary		Get Kaspa network analytics
//	@Description	Returns calculated analytics for the Kaspa network, including hashrate, supply, mining reward and next emission reduction.
//	@Tags			Analytics
//	@Produce		json
//	@Success		200	{object}	models.Analytics
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/analytics [get]
func AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	analytics, err := services.GetAnalytics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(analytics); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
