package middleware

import (
	"log/slog"

	"flypro-assessment/internal/config"

	"github.com/gin-gonic/gin"
)

func SetupMiddleware(router *gin.Engine, logger *slog.Logger, cfg config.Settings) {
	SetupCORS(router, cfg.Server.AllowedOrigins)
	cspMiddleware := CommonHeaders()
	router.Use(cspMiddleware)
	SetupRequestID(router)
	SetupRequestLogging(router, logger)
}
