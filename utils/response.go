package utils

import "github.com/gin-gonic/gin"

func ResponseData(data any) map[string]any {
	return gin.H{
		"data": data,
	}
}
