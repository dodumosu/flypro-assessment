package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS(router *gin.Engine, allowedOrigins []string) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allowedOrigins

	router.Use(cors.New(corsConfig))
}
