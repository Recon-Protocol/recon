package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Version string `json:"version"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Service: "recon-api",
		Version: "0.1.0",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("RECON API listening on :8081")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}