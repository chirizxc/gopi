package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func New(log *slog.Logger) gin.HandlerFunc {
	log = log.With(slog.String("component", "middleware/logger"))
	log.Info("logger middleware enabled")

	return func(c *gin.Context) {
		entry := log.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("remote_addr", c.ClientIP()),
			slog.String("user_agent", c.Request.UserAgent()),
		)

		start := time.Now()

		c.Next()

		duration := time.Since(start)

		entry.Info("request completed",
			slog.Int("status", c.Writer.Status()),
			slog.Int("bytes", c.Writer.Size()),
			slog.String("duration", duration.String()),
		)
	}
}
