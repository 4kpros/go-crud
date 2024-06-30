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

// @Tags Sign in
// @Summary Sign in user with email
// @Accept  json
// @Produce  json
// @Param   email body string true "Enter your email"
// @Param   password body string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 403 {string} string "Account not activated!"
// @Failure 404 {string} string "Invalid email or password!"
// @Router /auth/signin-email [post]
func (controller *AuthController) SignInWithEmail(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignInWithEmailRequest{}
	c.Bind(reqData)
	isEmailValid := utils.IsEmailValid(reqData.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
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

	// Execute the service
	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignInWithEmail(reqData)
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
		if errCode == http.StatusNotFound {
			tmpMessage := "User account not found! Please check your information."
			c.AbortWithError(errCode, fmt.Errorf("%s", tmpMessage))
			return
		}
		c.AbortWithError(errCode, err)
		return
	}

	// Return response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignInResponse{
			AccessToken: accessToken,
			Expires:     *accessExpires,
		},
	})
}

// @Tags Sign in
// @Summary Sign in user with phone number
// @Accept  json
// @Produce  json
// @Param   phoneNumber body int true "Enter your phone number"
// @Param   password body string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 403 {string} string "Account not activated!"
// @Failure 404 {string} string "Invalid phone number or password!"
// @Router /auth/signin-phone [post]
func (controller *AuthController) SignInWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignInWithPhoneNumberRequest{}
	c.Bind(reqData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
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

	// Execute the service
	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignInWithPhoneNumber(reqData)
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
		if errCode == http.StatusNotFound {
			tmpMessage := "User account not found! Please check your information."
			c.AbortWithError(errCode, fmt.Errorf("%s", tmpMessage))
			return
		}
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignInResponse{
			AccessToken: accessToken,
			Expires:     *accessExpires,
		},
	})
}

// @Tags Sign in
// @Summary Sign in user with provider
// @Accept  json
// @Produce  json
// @Param   provider body string true "Enter your provider" Enums(google, facebook, instagram)
// @Param   token body string true "Enter your token" minlength(8)
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 404 {string} string "Invalid token!"
// @Router /auth/signin-provider [post]
func (controller *AuthController) SignInWithProvider(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignInWithProviderRequest{}
	c.Bind(reqData)

	// Execute the service
	accessToken, accessExpires, errCode, err := controller.Service.SignInWithProvider(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignInResponse{
			AccessToken: accessToken,
			Expires:     *accessExpires,
		},
	})
}

// @Tags Sign up
// @Summary Sign up user with email
// @Accept  json
// @Produce  json
// @Param   email body string true "Enter your email"
// @Param   password body string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignUpResponse "OK"
// @Failure 302 {string} string "User with this email already exists!"
// @Router /auth/signup-email [post]
func (controller *AuthController) SignUpWithEmail(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignUpWithEmailRequest{}
	c.Bind(reqData)
	isEmailValid := utils.IsEmailValid(reqData.Email)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
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

	// Execute the service
	token, errCode, err := controller.Service.SignUpWithEmail(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignUpResponse{
			Token:   token,
			Message: "Account created! Please activate your account to start using your services.",
		},
	})
}

// @Tags Sign up
// @Summary Sign up user with phone number
// @Accept  json
// @Produce  json
// @Param   phoneNumber body int true "Enter your phone number"
// @Param   password body string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignUpResponse "OK"
// @Failure 302 {string} string "User with this phone number already exists!"
// @Router /auth/signup-phone [post]
func (controller *AuthController) SignUpWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var reqData = &request.SignUpWithPhoneNumberRequest{}
	c.Bind(reqData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
	isPasswordValid, missingPasswordChars := utils.IsPasswordValid(reqData.Password)
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

	// Execute the service
	token, errCode, err := controller.Service.SignUpWithPhoneNumber(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.SignUpResponse{
			Token:   token,
			Message: "Account created! Please activate your account to start using your services.",
		},
	})
}

// @Tags Activate account
// @Summary Activate new user account
// @Accept  json
// @Produce  json
// @Param   token body string true "Enter your token from sign in or sign up"
// @Param   code body int true "Enter your received code"
// @Success 200 {object} response.ActivateAccountResponse "OK"
// @Failure 400 {string} string "Invalid token!"
// @Failure 403 {string} string "User account is already activated!"
// @Failure 404 {string} string "User not found!"
// @Router /auth/activate [post]
func (controller *AuthController) ActivateAccount(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ActivateAccountRequest{}
	c.Bind(reqData)

	// Execute the service
	activatedAt, errCode, err := controller.Service.ActivateAccount(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ActivateAccountResponse{
			ActivatedAt: *activatedAt,
		},
	})
}

// @Tags Reset password
// @Summary Reset password  with email init
// @Accept  json
// @Produce  json
// @Param   email body string true "Enter your email"
// @Success 200 {object} response.ResetPasswordInitResponse "OK"
// @Failure 400 {string} string "Invalid email input!"
// @Failure 404 {string} string "User with email not found!"
// @Router /auth/reset/init-email [post]
func (controller *AuthController) ResetPasswordEmailInit(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordEmailInitRequest{}
	c.Bind(reqData)
	isEmailValid := utils.IsEmailValid(reqData.Email)
	if !isEmailValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.ResetPasswordEmailInit(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		message := "Failed to start the process! Please try again later."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordInitResponse{
			Token: token,
		},
	})
}

// @Tags Reset password
// @Summary Reset password with phone number init
// @Accept  json
// @Produce  json
// @Param   phoneNumber body int true "Enter your phone number"
// @Success 200 {object} response.ResetPasswordInitResponse "OK"
// @Failure 400 {string} string "Invalid phone number input!"
// @Failure 404 {string} string "User with phone number not found!"
// @Router /auth/reset/init-phone [post]
func (controller *AuthController) ResetPasswordPhoneNumberInit(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordPhoneNumberInitRequest{}
	c.Bind(reqData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(reqData.PhoneNumber)
	if !isPhoneNumberValid {
		message := "Invalid phone number! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Execute the service
	token, errCode, err := controller.Service.ResetPasswordPhoneNumberInit(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}
	if len(token) <= 0 {
		errCode = http.StatusInternalServerError
		message := "Failed to start the process! Please try again later."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordInitResponse{
			Token: token,
		},
	})
}

// @Tags Reset password
// @Summary Reset password set code
// @Accept  json
// @Produce  json
// @Param   token body string true "Enter your token"
// @Param   code body int true "Enter your received code"
// @Success 200 {object} response.ResetPasswordInitResponse "OK"
// @Failure 400 {string} string "Invalid token!"
// @Failure 404 {string} string "User not found!"
// @Router /auth/reset/code [post]
func (controller *AuthController) ResetPasswordCode(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordCodeRequest{}
	c.Bind(reqData)

	// Execute the service
	token, errCode, err := controller.Service.ResetPasswordCode(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordCodeResponse{
			Token: token,
		},
	})
}

// @Tags Reset password
// @Summary Reset password set new password
// @Accept  json
// @Produce  json
// @Param   token body string true "Enter your token"
// @Param   newPassword body string true "Enter your new password"
// @Success 200 {object} response.ResetPasswordNewPasswordResponse "OK"
// @Failure 400 {string} string "Invalid token or password input!"
// @Failure 404 {string} string "User not found!"
// @Router /auth/reset/new-password [post]
func (controller *AuthController) ResetPasswordNewPassword(c *gin.Context) {
	// Get data of req body
	var reqData = &request.ResetPasswordNewPasswordRequest{}
	c.Bind(reqData)

	// Execute the service
	errCode, err := controller.Service.ResetPasswordNewPassword(reqData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.ResetPasswordNewPasswordResponse{
			Message: "Password successful changed! Please sign in to start using our services.",
		},
	})
}
