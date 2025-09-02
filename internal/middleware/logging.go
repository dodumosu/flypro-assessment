package middleware

import (
	"log/slog"

	"github.com/gin-contrib/requestid"
	slogger "github.com/gin-contrib/slog"
	"github.com/gin-gonic/gin"
)

func SetupRequestLogging(r *gin.Engine, logger *slog.Logger) {
	r.Use(slogger.SetLogger(
		slogger.WithLogger(func(ctx *gin.Context, l *slog.Logger) *slog.Logger {
			return logger.With("request_id", requestid.Get(ctx))
		}),
	))
}
