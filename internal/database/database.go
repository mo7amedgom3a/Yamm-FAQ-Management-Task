package database

import (
	"fmt"
	"log"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBPassword, cfg.DBName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database successfully")
}
