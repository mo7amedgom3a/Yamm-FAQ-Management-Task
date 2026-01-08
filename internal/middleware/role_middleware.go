package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get(ContextUserRoleKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user role not found in context"})
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid role type"})
			return
		}

		for _, role := range allowedRoles {
			if role == roleStr {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient privileges"})
	}
}
