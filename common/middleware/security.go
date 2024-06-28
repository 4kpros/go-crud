package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/gin-gonic/gin"
)

func CheckIfWeCanTrustOrigin(c *gin.Context) (trust bool) {
	host := c.Request.Host
	fmt.Printf("\nHTTP request from host ==> %s\n", c.Request.Host)
	if strings.EqualFold(host, "localhost:3000") || strings.EqualFold(host, "127.0.0.1:3000") {
		trust = true
		return
	}
	trust = false
	return
}

func SecureAPIHandler(handler gin.HandlerFunc, requiredAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		canTrust := CheckIfWeCanTrustOrigin(c)
		if !canTrust {
			message := "Our system detected your request as malicious! Please fix that before."
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Xss-Protection", "1; mode=block")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("X-Download-Options", "noopen")
		c.Header("Strict-Transport-Security", fmt.Sprintf("max-age=%d; %s", 31536000, "includeSubDomains"))
		c.Next()
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != config.AppEnv.ApiKey {
			message := "Invalid API key! Please enter valid API key and try again."
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
		} else {
			if requiredAuth {
				JWTHandler(c, handler)
			} else {
				handler(c)
			}
		}
	}
}

func JWTHandler(c *gin.Context, handler gin.HandlerFunc) {
	bearerToken := c.GetHeader("Authorization")
	if len(bearerToken) <= 0 {
		message := "Missing authorization header! Please enter authorization header and try again."
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", message))
		return
	}
	if !utils.VerifyJWTToken(strings.TrimPrefix(bearerToken, "Bearer ")) {
		message := "Invalid authorization header! Please enter valid authorization header and try again."
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", message))
		return
	}
	handler(c)
}
