package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// HashrateHandler godoc
//
//	@Summary		Get Kaspa network hashrate
//	@Description	Returns the current Kaspa network hashrate.
//	@Tags			Hashrate
//	@Produce		json
//	@Success		200	{object}	models.HashrateInfo
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/hashrate [get]
func HashrateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hashrate, err := services.GetHashrate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(hashrate); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
