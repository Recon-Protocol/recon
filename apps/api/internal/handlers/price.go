package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// PriceHandler godoc
//
//	@Summary		Get Kaspa price information
//	@Description	Returns the current Kaspa price, market capitalization, 24-hour volume and 24-hour price change.
//	@Tags			Price
//	@Produce		json
//	@Success		200	{object}	models.PriceInfo
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/api/v1/price [get]
func PriceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	price, err := services.GetPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(price); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
