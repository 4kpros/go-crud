package controllers

import (
	"net/http"

	"github.com/4kpros/go-crud/common/utils"
	"github.com/gin-gonic/gin"
)

func ResetPasswordInit(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func ResetPasswordCode(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func ResetPasswordSecretQuestion(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func ResetPasswordNewPassword(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}
