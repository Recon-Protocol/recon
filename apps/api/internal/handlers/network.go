package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// NetworkHandler godoc
//
//	@Summary		Get Kaspa network information
//	@Description	Returns the current Kaspa network information, including block count, header count, difficulty and virtual DAA score.
//	@Tags			Network
//	@Produce		json
//	@Success		200	{object}	models.NetworkInfo
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/network [get]
func NetworkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	network, err := services.GetNetwork()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(network); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
