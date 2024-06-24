package middleware

import (
	"github.com/4kpros/go-crud/errors"
	"github.com/gin-gonic/gin"
)

func ErrorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			c.AbortWithStatusJSON(errors.APIError(
				c.Writer.Status(),
				err.Err.Error(),
			))
			// if err.Err.Error() != "EOF" {
			// 	c.AbortWithStatusJSON(errors.APIError(
			// 		c.Writer.Status(),
			// 		err.Err.Error(),
			// 	))
			// }
		}
	}
}
