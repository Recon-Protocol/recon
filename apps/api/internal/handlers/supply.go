package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// SupplyHandler godoc
//
//	@Summary		Get Kaspa supply information
//	@Description	Returns the current Kaspa supply and emission data.
//	@Tags			Supply
//	@Produce		json
//	@Success 200 {object} models.SupplyInfo
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/supply [get]
func SupplyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	supply, err := services.GetSupply()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(supply); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
