package scripts

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/database"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)

	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	adminEmail := cfg.AdminEmail
	adminPassword := cfg.AdminPassword

	var existingUser models.User
	if err := database.DB.Where("email = ?", adminEmail).First(&existingUser).Error; err == nil {
		log.Println("Admin user already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}

	admin := models.User{
		ID:           uuid.New(),
		Email:        adminEmail,
		PasswordHash: string(hashedPassword),
		Role:         models.RoleAdmin,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		log.Fatal("Failed to create admin user: ", err)
	}

	log.Println("Admin user created successfully")
}
