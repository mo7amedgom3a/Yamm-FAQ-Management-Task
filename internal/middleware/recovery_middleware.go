package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().
					Interface("error", err).
					Bytes("stack", debug.Stack()).
					Msg("Panic recovered")

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}
