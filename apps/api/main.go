package main

import (
	"github.com/Recon-Protocol/recon/apps/api/internal/config"
	"github.com/Recon-Protocol/recon/apps/api/internal/router"
	"github.com/Recon-Protocol/recon/apps/api/internal/server"
)

func main() {
	cfg := config.Load()

	r := router.New()

	server.Start(cfg, r)
}
