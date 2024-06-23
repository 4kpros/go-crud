package middleware

import (
	"net/http"

	"github.com/4kpros/go-crud/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for status, err := range c.Errors {
			if status > 0 {
				c.AbortWithStatusJSON(APIError(
					status,
					err.Err.Error(),
				))
			} else {
				c.AbortWithStatusJSON(APIError(
					http.StatusInternalServerError,
					"Something happens! Please try again later.",
				))
			}
		}
	}
}

func APIError(statusCode int, message string) (code int, obj any) {
	statusText := http.StatusText(statusCode)
	utils.Logger.Warn(
		message,
		zap.Int("status", statusCode),
		zap.String("statusText", statusText),
	)
	return statusCode, gin.H{
		"StatusCode": statusCode,
		"StatusText": statusText,
		"Message":    message,
	}
}
