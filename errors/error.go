package errors

import (
	"net/http"

	"github.com/4kpros/go-crud/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
