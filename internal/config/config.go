package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName            string
	DBHost            string
	DBPassword        string
	AdminEmail        string
	AdminPassword     string
	JWTSecret         string
	JWTExpirationTime int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	jwtExpirationStr := os.Getenv("JWT_EXPIRATION_TIME")
	jwtExpiration, _ := strconv.Atoi(jwtExpirationStr)

	return &Config{
		DBName:            os.Getenv("DB_NAME"),
		DBHost:            os.Getenv("DB_HOST"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		AdminEmail:        os.Getenv("ADMIN_EMAIL"),
		AdminPassword:     os.Getenv("ADMIN_PASSWORD"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTExpirationTime: jwtExpiration,
	}
}
