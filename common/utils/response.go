package utils

import (
	"github.com/4kpros/go-crud/common/types"
	"github.com/gin-gonic/gin"
)

func ResponseDataWithPagination(data any, pagination types.Pagination, filter types.Filter) map[string]any {
	return gin.H{
		"Data":       data,
		"Pagination": pagination,
		"Filter":     filter,
	}
}

func ResponseData(data any) map[string]any {
	return gin.H{
		"Data": data,
	}
}
