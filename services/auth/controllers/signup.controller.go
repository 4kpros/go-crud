package controllers

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-crud/common/types"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/config"
	"github.com/4kpros/go-crud/services/auth/data"
	"github.com/4kpros/go-crud/services/auth/models"
	"github.com/gin-gonic/gin"
)

func SignUpWithEmail(c *gin.Context) {
	// Get data of req body
	var requestData data.SignUpWithEmailRequest
	c.Bind(&requestData)
	isValidEmail := utils.IsEmailValid(requestData.Email)
	isValidPassword, missingPasswordChars := utils.IsPasswordValid(requestData.Password)
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
	var newUser models.NewUser
	config.DB.Where("email = ? AND (provider is null OR provider = '')", requestData.Email).Limit(1).Find(&newUser)
	if utils.IsEmailValid(newUser.Email) {
		message := "User with same email already exists! Please use another email address."
		c.AbortWithError(http.StatusFound, fmt.Errorf("%s", message))
		return
	}

	// Create a user
	newUser.Email = requestData.Email
	newUser.Password = requestData.Password
	result := config.DB.Create(&newUser)
	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	// Create JWT Token
	//

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: newUser,
	})
}

func SignUpWithPhoneNumber(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}

func SignUpWithGoogle(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}

func SignUpWithFacebook(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: nil,
	})
}
