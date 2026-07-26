package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

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
