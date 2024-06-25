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

	// Check if user exists
	var existingNewUser models.NewUser
	var encodedPassword = utils.EncryptPassword(dataReq.Password)
	initializers.DB.Where("email = ? AND password = ?", dataReq.Email, encodedPassword).Limit(1).Find(&existingNewUser)
	if !utils.IsEmailValid(existingNewUser.Email) {
		message := "Invalid email address or password! Please enter valid information."
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("%s", message))
		return
	}

	// Return resp
	existingNewUser.Password = ""
	c.JSON(http.StatusOK, utils.ResponseData(existingNewUser))
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
