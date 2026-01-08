package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	HeaderXRequestID = "X-Request-ID"
	ContextRequestID = "requestID"
)

func ContextMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request ID
		requestID := c.GetHeader(HeaderXRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(ContextRequestID, requestID)
		c.Header(HeaderXRequestID, requestID)

		// Timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
