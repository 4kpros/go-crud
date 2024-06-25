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

func SignInWithEmail(c *gin.Context) {
	// Get data of req body
	var requestData data.SignInWithEmailRequest
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

	// Check if user exists
	var newUser models.NewUser
	var encryptedPassword, _ = utils.EncryptWithArgon2id(requestData.Password)
	config.DB.Where("email = ? AND password = ?", newUser.Email, encryptedPassword).Limit(1).Find(&newUser)
	if !utils.IsEmailValid(newUser.Email) {
		message := "Invalid email address or password! Please enter valid information."
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("%s", message))
		return
	}

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: data.SignInResponse{},
	})
}

func SignInWithPhoneNumber(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: data.SignInResponse{},
	})
}

func SignInWithGoogle(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: data.SignInResponse{},
	})
}

func SignInWithFacebook(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: data.SignInResponse{},
	})
}

func SignInWith2fa(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: data.SignInResponse{},
	})
}
