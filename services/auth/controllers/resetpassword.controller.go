package controllers

import (
	"net/http"

	"github.com/4kpros/go-crud/common/types"
	"github.com/gin-gonic/gin"
)

func ResetPasswordInit(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}

func ResetPasswordCode(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}

func ResetPasswordSecretQuestion(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}

func ResetPasswordNewPassword(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}
