package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SecureAPIKeyHandler(handler gin.HandlerFunc, requiredAuth bool) gin.HandlerFunc {
	var effectiveApiKey = os.Getenv("API_KEY")
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != effectiveApiKey {
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
		message := "You need to login before accessing this resource!"
		c.AbortWithError(http.StatusForbidden, fmt.Errorf("%s", message))
	} else {
		handler(c)
	}
}
