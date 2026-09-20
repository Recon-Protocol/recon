package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// ReportHandler godoc
//
//	@Summary		Get Kaspa network report
//	@Description	Returns a formatted report containing the current Kaspa network, supply, hashrate and emission data.
//	@Tags			Report
//	@Produce		json
//	@Success		200	{object}	models.Report
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/report [get]
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	report, err := services.GetReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(report); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
