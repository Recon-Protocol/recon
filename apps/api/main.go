package main

import (
	"github.com/Recon-Protocol/recon/apps/api/internal/config"
	"github.com/Recon-Protocol/recon/apps/api/internal/router"
	"github.com/Recon-Protocol/recon/apps/api/internal/server"
)

// @title RECON API
// @version 0.1.0
// @description RECON API for Kaspa network analytics and intelligence.
// @contact.name Recon Protocol
// @license.name MIT
// @BasePath /api/v1
func main() {
	cfg := config.Load()

	r := router.New()

	server.Start(cfg, r)
}
