package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// MetricsHandler godoc
//
//	@Summary		Get Kaspa network metrics
//	@Description	Returns the current Kaspa network metrics snapshot, including block count, header count, difficulty and virtual DAA score.
//	@Tags			Metrics
//	@Produce		json
//	@Success		200	{object}	models.NetworkSnapshot
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/metrics [get]
func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	snapshot, err := services.GetSnapshot()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(snapshot); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
