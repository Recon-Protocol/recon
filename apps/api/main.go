package main

import (
	"log"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/config"
	"github.com/Recon-Protocol/recon/apps/api/internal/handlers"
)

func main() {

	cfg := config.Load()

	http.HandleFunc("/health", handlers.HealthHandler)

	log.Printf("RECON API listening on :%s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}