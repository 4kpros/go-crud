package controllers

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-crud/common/initializers"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/services/auth/models"
	"github.com/gin-gonic/gin"
)

func SignUpWithEmail(c *gin.Context) {
	// Get data of req body
	var dataReq struct {
		Email    string
		Password string
	}
	c.Bind(&dataReq)
	isValidEmail := utils.IsEmailValid(dataReq.Email)
	isValidPassword, missingPasswordChars := utils.IsPasswordValid(dataReq.Password)
	if !isValidEmail && !isValidPassword {
		message := "Invalid email and password! Please enter valid email address and password. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isValidEmail {
		message := "Invalid email! Please enter valid email address."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isValidPassword {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Check if user doesn't exists
	var existingNewUser models.NewUser
	initializers.DB.Where("email = ?", dataReq.Email).Limit(1).Find(&existingNewUser)
	if utils.IsEmailValid(existingNewUser.Email) {
		message := "User with same email already exists ! Please use another email address."
		c.AbortWithError(http.StatusFound, fmt.Errorf("%s", message))
		return
	}

	// Create a user
	result := initializers.DB.Create(&dataReq)
	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	// Return resp
	c.JSON(http.StatusOK, utils.ResponseData(dataReq))
}

func SignUpWithPhoneNumber(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func SignUpWithGoogle(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func SignUpWithFacebook(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func ValidateAccount(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}
