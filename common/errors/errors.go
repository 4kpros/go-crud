package errors

import (
	"net/http"

	"github.com/4kpros/go-crud/common/types"
	"github.com/4kpros/go-crud/common/utils"
	"go.uber.org/zap"
)

func APIError(statusCode int, message string) (int, interface{}) {
	statusText := http.StatusText(statusCode)
	utils.Logger.Warn(
		message,
		zap.Int("status", statusCode),
		zap.String("statusText", statusText),
	)

	return statusCode, types.WebErrorResponse{
		Code:    statusCode,
		Status:  statusText,
		Message: message,
	}
}
