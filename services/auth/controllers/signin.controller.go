package controllers

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-crud/common/initializers"
	"github.com/4kpros/go-crud/common/utils"
	"github.com/4kpros/go-crud/services/auth/models"
	"github.com/gin-gonic/gin"
)

func SignInWithEmail(c *gin.Context) {
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

	// Check if user exists
	var existingNewUser models.NewUser
	initializers.DB.Where("email = ? AND password = ?", dataReq.Email, dataReq.Password).Limit(1).Find(&existingNewUser)
	if !utils.IsEmailValid(existingNewUser.Email) {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("Invalid email address or password ! Please enter valid information."))
		return
	}

	// Return resp
	c.JSON(http.StatusOK, utils.ResponseData(dataReq))
}

func SignInWithPhoneNumber(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func SignInWithGoogle(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func SignInWithFacebook(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}

func SignInWith2fa(c *gin.Context) {
	c.JSON(http.StatusOK, utils.ResponseData(nil))
}
