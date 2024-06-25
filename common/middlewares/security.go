package middlewares

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/gin-gonic/gin"
)

func SecureAPIKeyHandler(handler gin.HandlerFunc, requiredAuth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != config.EnvConfig.ApiKey {
			message := "Invalid API key! Please enter valid API key and try again."
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
		} else {
			if requiredAuth {
				SecureJWTHandler(c, handler)
			} else {
				handler(c)
			}
		}
	}
}

func SecureJWTHandler(c *gin.Context, handler gin.HandlerFunc) {
	bearerToken := c.GetHeader("Authorization")
	if len(bearerToken) <= 0 {
		message := "Missing authorization header! Please enter authorization header and try again."
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", message))
		return
	}
	if !utils.VerifyJWTToken(bearerToken) {
		message := "Invalid authorization header! Please enter valid authorization header and try again."
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("%s", message))
		return
	}
	handler(c)
}
