package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

func GenerateToken(user *models.User, cfg *config.Config) (string, error) {
	JWTKey := []byte(cfg.JWTSecret)
	expirationTime := time.Duration(cfg.JWTExpirationTime) * time.Second

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(expirationTime).Unix(),
	})

	return token.SignedString(JWTKey)
}

func VerifyToken(tokenString string, cfg *config.Config) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}