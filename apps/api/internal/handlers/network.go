package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

func NetworkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := services.GetNetwork()

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}