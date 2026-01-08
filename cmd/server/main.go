package main

import (
	"context"
	"fmt"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/auth"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/database"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
)

func main() {
	// fmt.Println("seed admin user")
	// scripts.SeedAdmin()

	cfg := config.LoadConfig()
	database.ConnectDB(cfg)
	userRepo := repositories.NewUserRepository(database.DB)

	// test user reposito
	user := models.User{
		Email:        "test6@gmail.com",
		PasswordHash: "test-password",
		Role:         "merchant",
	}

	err := userRepo.CreateUser(context.Background(), &user)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("user", user)
	token, err := auth.GenerateToken(&user, cfg)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated token:", token)

	err = auth.VerifyToken(token, cfg)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}
	fmt.Println("Token verified successfully")
	fmt.Println("Extracted claims ...")
	claims, err := auth.ExtractClaims(token, cfg)
	if err != nil {
		fmt.Println("Error extracting claims:", err)
		return
	}
	fmt.Println("Extracted claims:", claims["user_id"])

}
