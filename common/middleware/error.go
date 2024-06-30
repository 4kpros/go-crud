package middleware

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/gin-gonic/gin"
)

func ErrorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			if err.Err.Error() != "EOF" {
				c.AbortWithStatusJSON(c.Writer.Status(), types.WebErrorResponse{
					Message: err.Err.Error(),
				})
			}
		}
	}
}
