package controllers

import (
	"net/http"

	"github.com/4kpros/go-crud/common/types"
	"github.com/gin-gonic/gin"
)

func ValidateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}
