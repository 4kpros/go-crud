package controller

import (
	"fmt"
	"net/http"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/features/auth/data/request"
	"github.com/4kpros/go-api/features/auth/data/response"
	"github.com/4kpros/go-api/features/auth/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (controller *AuthController) SignInWithEmail(c *gin.Context) {
	// Get data of req body
	var requestData request.SignInWithEmailRequest
	c.Bind(&requestData)
	isEmailValid := utils.IsEmailValid(requestData.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(requestData.Password)
	if !isEmailValid && !isPasswordValid {
		message := "Invalid email and password! Please enter valid information. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isEmailValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Generate token
	validateAccountToken, accessToken, refreshToken, accessExpires, refreshExpires, errCode, err := controller.Service.SignInWithEmail(&requestData)
	if err != nil {
		if errCode == http.StatusForbidden || len(validateAccountToken) > 0 {
			c.JSON(http.StatusForbidden, types.WebSuccessResponse{
				Data: response.SignUpResponse{
					Token:   validateAccountToken,
					Message: "Account not activated! Please activate your account to start using your services.",
				},
			})
			return
		}
		c.AbortWithError(errCode, err)
		return
	}

	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignInResponse{
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			AccessExpires:  accessExpires,
			RefreshExpires: refreshExpires,
		},
	})
}

func (controller *AuthController) SignInWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var requestData request.SignInWithPhoneNumberRequest
	c.Bind(&requestData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(requestData.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(requestData.Password)
	if !isPhoneNumberValid && !isPasswordValid {
		message := "Invalid phone number and password! Please enter valid information. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPhoneNumberValid {
		message := "Invalid phone number! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Generate token
	validateAccountToken, accessToken, refreshToken, accessExpires, refreshExpires, errCode, err := controller.Service.SignInWithPhoneNumber(&requestData)
	if err != nil {
		if errCode == http.StatusForbidden || len(validateAccountToken) > 0 {
			c.JSON(http.StatusForbidden, types.WebSuccessResponse{
				Data: response.SignUpResponse{
					Token:   validateAccountToken,
					Message: "Account not activated! Please activate your account to start using your services.",
				},
			})
			return
		}
		c.AbortWithError(errCode, err)
		return
	}

	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignInResponse{
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			AccessExpires:  accessExpires,
			RefreshExpires: refreshExpires,
		},
	})
}

func (controller *AuthController) SignInWithProvider(c *gin.Context) {
	// Get data of req body
	var requestData request.SignInWithProviderRequest
	c.Bind(&requestData)

	// Generate token
	accessToken, refreshToken, accessExpires, refreshExpires, errCode, err := controller.Service.SignInWithProvider(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignInResponse{
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			AccessExpires:  accessExpires,
			RefreshExpires: refreshExpires,
		},
	})
}

func (controller *AuthController) SignUpWithEmail(c *gin.Context) {
	// Get data of req body
	var requestData request.SignUpWithEmailRequest
	c.Bind(&requestData)
	isEmailValid := utils.IsEmailValid(requestData.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(requestData.Password)
	if !isEmailValid && !isPasswordValid {
		message := "Invalid email and password! Please enter valid email address and password. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isEmailValid {
		message := "Invalid email! Please enter valid email address."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Generate account validation token
	token, errCode, err := controller.Service.SignUpWithEmail(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignUpResponse{
			Token:   token,
			Message: "Account created! Please activate your account to start using your services.",
		},
	})
}

func (controller *AuthController) SignUpWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var requestData request.SignUpWithPhoneNumberRequest
	c.Bind(&requestData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(requestData.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(requestData.Password)
	if !isPhoneNumberValid && !isPasswordValid {
		message := "Invalid phone number and password! Please enter valid phone number and password. Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPhoneNumberValid {
		message := "Invalid email! Please enter valid phone number."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPasswordValid {
		message := "Invalid password! Password missing " + missingPasswordChars
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Generate account validation token
	token, errCode, err := controller.Service.SignUpWithPhoneNumber(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignUpResponse{
			Token:   token,
			Message: "Account created! Please activate your account to start using your services.",
		},
	})
}

func (controller *AuthController) ActivateAccount(c *gin.Context) {
	// Get data of req body
	var requestData request.ActivateAccountRequest
	c.Bind(&requestData)

	// Generate account validation token
	activatedAt, errCode, err := controller.Service.ActivateAccount(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ActivateAccountResponse{
			ActivatedAt: *activatedAt,
		},
	})
}

func (controller *AuthController) ResetPasswordInit(c *gin.Context) {
	reqData := &request.ResetPasswordInitRequest{}
	token, errCode, err := controller.Service.ResetPasswordInit(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
	}
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordInitResponse{
			Token: token,
		},
	})
}

func (controller *AuthController) ResetPasswordCode(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordCodeResponse{
			Token: "",
		},
	})
}

func (controller *AuthController) ResetPasswordNewPassword(c *gin.Context) {
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordNewPasswordResponse{
			Message: "",
		},
	})
}
