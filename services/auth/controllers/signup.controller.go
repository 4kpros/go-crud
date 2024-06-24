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
	var dataReq models.NewUser
	c.Bind(&dataReq)
	isValidEmail := utils.IsEmailValid(dataReq.Email)
	isValidPassword, missingPasswordChars := utils.IsPasswordValid(dataReq.Password)
	if !isValidEmail && !isValidPassword {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid email and password ! Please enter valid email address and password. Password missing %s", missingPasswordChars))
		return
	}
	if !isValidEmail {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid email ! Please enter valid email address."))
		return
	}
	if !isValidPassword {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid password ! Password missing %s", missingPasswordChars))
		return
	}

	// Check if user doesn't exists
	var existingNewUser models.NewUser
	initializers.DB.Where("email = ?", dataReq.Email).Limit(1).Find(&existingNewUser)
	if utils.IsEmailValid(existingNewUser.Email) {
		c.AbortWithError(http.StatusFound, fmt.Errorf("User with same email already exists ! Please use another email address."))
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
