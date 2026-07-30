package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

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
