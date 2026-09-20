package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

func IntelligenceHandler(w http.ResponseWriter, r *http.Request) {
	intel, err := services.GetIntelligence()
	if err != nil {
		http.Error(w, "failed to compute intelligence", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(intel); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
