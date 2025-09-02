package middleware

import (
	"flypro-assessment/internal/config"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func SetupMiddleware(router *gin.Engine, logger *slog.Logger, cfg config.Settings) {
	cspMiddleware := CommonHeaders()
	router.Use(cspMiddleware)
	SetupRequestID(router)
	SetupRequestLogging(router, logger)
	SetupCORS(router, cfg.Server.AllowedOrigins)
}
