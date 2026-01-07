package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

func GenerateToken(user *models.User, cfg *config.Config) (string, error) {
	var err error
	JWTKey := []byte(cfg.JWTSecret)
	expirationTime := cfg.JWTExpirationTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Duration(expirationTime)).Unix(),
	})
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}	

func VerifyToken(tokenString string, cfg *config.Config) error {
	var err error
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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