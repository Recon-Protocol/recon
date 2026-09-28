package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/services"
)

// BlockDAGHandler godoc
//
//	@Summary		BlockDAG Live Tips
//	@Description	Returns current BlockDAG tip blocks with DAA score and tx count.
//	@Tags			Network
//	@Produce		json
//	@Success		200	{object}	models.BlockDAG
//	@Router			/api/v1/blockdag [get]
func BlockDAGHandler(w http.ResponseWriter, r *http.Request) {
	blockdag, err := services.GetBlockDAG()
	if err != nil {
		http.Error(w, "failed to fetch blockdag", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(blockdag); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}