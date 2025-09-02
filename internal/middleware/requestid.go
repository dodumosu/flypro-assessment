package middleware

import (
	"flypro-assessment/internal/utils"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func SetupRequestID(r *gin.Engine) {
	r.Use(requestid.New(
		requestid.WithGenerator(func() string {
			return utils.NewID()
		}),
	))
}
