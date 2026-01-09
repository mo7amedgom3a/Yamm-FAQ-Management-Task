package main

import (
	"fmt"
	"log"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/database"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/routes"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/scripts"
)

func main() {
	scripts.SeedAdmin()
	cfg := config.LoadConfig()

	database.ConnectDB(cfg)
	db := database.DB

	r := routes.SetupRouter(db, cfg)

	port := cfg.ServerPort
	if cfg.DebugMode == "debug" {
		fmt.Printf("Server starting on port %v...\n", port)
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
