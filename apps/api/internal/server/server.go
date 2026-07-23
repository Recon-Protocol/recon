package server

import (
	"log"
	"net/http"

	"github.com/Recon-Protocol/recon/apps/api/internal/config"
)

func Start(cfg config.Config, handler http.Handler) {
	log.Printf("RECON API listening on :%s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}