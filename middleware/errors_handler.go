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
		for _, err := range c.Errors {
			if err.Err.Error() != "EOF" {
				c.AbortWithStatusJSON(APIError(
					c.Writer.Status(),
					err.Err.Error(),
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
