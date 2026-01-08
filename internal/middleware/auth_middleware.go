package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/auth"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
)

const (
	ContextUserIDKey   = "userID"
	ContextUserRoleKey = "userRole"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		err := auth.VerifyToken(tokenString, cfg)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, err := auth.ExtractClaims(tokenString, cfg)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "failed to extract claims"})
			return
		}

		c.Set(ContextUserIDKey, claims["user_id"])
		c.Set(ContextUserRoleKey, claims["role"])
		c.Next()
	}
}
