package main

import (
	"fmt"
	"log"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/database"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/routes"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	// Connect to Database
	database.ConnectDB(cfg)
	db := database.DB

	// Setup Router
	r := routes.SetupRouter(db, cfg)

	// Run Server
	port := cfg.ServerPort
	if cfg.DebugMode == "debug" {
		fmt.Printf("Server starting on port %v...\n", port)
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
