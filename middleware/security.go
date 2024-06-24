package middleware

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
			c.AbortWithError(
				http.StatusForbidden,
				fmt.Errorf("invalid API-KEY, Please enter valid API key and try again"),
			)
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
		c.AbortWithError(
			http.StatusUnauthorized,
			fmt.Errorf("you need to login before to access this resource"),
		)
	} else {
		handler(c)
	}
}
