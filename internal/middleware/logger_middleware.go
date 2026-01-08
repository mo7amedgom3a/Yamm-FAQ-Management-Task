package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if raw != "" {
			path = path + "?" + raw
		}

		msg := "Request"
		if len(c.Errors) > 0 {
			msg = c.Errors.String()
		}

		logger := log.Info()
		if c.Writer.Status() >= 400 {
			logger = log.Error()
		}

		logger.
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Str("user_agent", c.Request.UserAgent()).
			Msg(msg)
	}
}
