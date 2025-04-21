package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryota1119/time_resport_webapi/internal/infrastructure/logger"
)

func Default() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		loggerMiddleware(),
		corsMiddleware(),
	}
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()

		reqLogger := logger.Logger.With(
			slog.String("request_id", requestID),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)

		ctx := logger.WithContext(c.Request.Context(), reqLogger)
		c.Request = c.Request.WithContext(ctx)

		reqLogger.Info("request started")

		c.Next()

		reqLogger.Info(
			"request complete",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Duration("duration", time.Since(start)),
		)
	}
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}
