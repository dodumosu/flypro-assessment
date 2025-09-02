package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CommonHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cspDirectives = []string{
			"default-src 'self'",
			"script-src 'self' https://unpkg.com", // For Huma docs UI
			"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com https://unpkg.com", // For Huma docs UI & fonts
			"connect-src 'self' https://unpkg.com",                                            // For Huma docs UI
			"img-src 'self' data: https://unpkg.com",                                          // For Huma docs UI
			"font-src 'self' https://fonts.gstatic.com",
			"worker-src 'self' blob:", // Huma docs UI might use workers
			"form-action 'self'",
			"frame-ancestors 'none'",
		}
		c.Header("Content-Security-Policy", strings.Join(cspDirectives, "; "))
		c.Header("Referrer-Policy", "origin-when-cross-origin") // A common, reasonable default
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")                                          // Prevent clickjacking
		c.Header("X-XSS-Protection", "0")                                            // Modern browsers have better built-in XSS protection; CSP is preferred.
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains") // If serving over HTTPS

		c.Next()
	}
}
