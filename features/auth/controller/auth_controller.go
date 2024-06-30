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
// @Param   email path string true "Enter your email"
// @Param   password path string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 403 {string} string "Account not activated!"
// @Failure 404 {string} string "Invalid email or password!"
// @Router /signin-email [post]
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
	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignInWithEmail(&requestData)
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
// @Param   phoneNumber path string true "Enter your phone number"
// @Param   password path string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 403 {string} string "Account not activated!"
// @Failure 404 {string} string "Invalid phone number or password!"
// @Router /signin-phone [post]
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
	validateAccountToken, accessToken, accessExpires, errCode, err := controller.Service.SignInWithPhoneNumber(&requestData)
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
// @Param   provider path string true "Enter your provider" Enums(google, facebook, instagram)
// @Param   token path string true "Enter your token" minlength(8)
// @Success 200 {object} response.SignInResponse "OK"
// @Failure 400 {string} string "Invalid inputs!"
// @Failure 404 {string} string "Invalid token!"
// @Router /signin-provider [post]
func (controller *AuthController) SignInWithProvider(c *gin.Context) {
	// Get data of req body
	var requestData request.SignInWithProviderRequest
	c.Bind(&requestData)

	// Generate token
	accessToken, accessExpires, errCode, err := controller.Service.SignInWithProvider(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

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
// @Param   email path string true "Enter your email"
// @Param   password path string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignUpResponse "OK"
// @Failure 302 {string} string "User with this email already exists!"
// @Router /signup-email [post]
func (controller *AuthController) SignUpWithEmail(c *gin.Context) {
	// Get data of req body
	var requestData request.SignUpWithEmailRequest
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

// @Tags Sign up
// @Summary Sign up user with phone number
// @Accept  json
// @Produce  json
// @Param   phoneNumber path string true "Enter your phone number"
// @Param   password path string true "Must be at least 8 characters with 1 upper case, 1 lower case, 1 special character and 1 number" minlength(8)
// @Success 200 {object} response.SignUpResponse "OK"
// @Failure 302 {string} string "User with this phone number already exists!"
// @Router /signup-phone [post]
func (controller *AuthController) SignUpWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var requestData request.SignUpWithPhoneNumberRequest
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

// @Tags Add new user - [super-admin]
// @Summary Add new user with email
// @Accept  json
// @Produce  json
// @Param   email path string true "Enter your email"
// @Param   role path string true "Select role" Enums(super-admin, admin, manager, manager-assist, deliver, customer, customer-service)
// @Success 200 {object} response.NewUserWithEmailResponse "OK"
// @Failure 400 {string} string "Invalid email or role!"
// @Failure 302 {string} string "User with this email already exists!"
// @Router /add-new-user-with-email [post]
func (controller *AuthController) NewUserWithEmail(c *gin.Context) {
	// Get data of req body
	var requestData request.NewUserWithEmailRequest
	c.Bind(&requestData)
	isEmailValid := utils.IsEmailValid(requestData.Email)
	isRoleValid := utils.IsRoleValid(requestData.Role)
	if !isEmailValid && !isRoleValid {
		message := "Invalid email and role! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isEmailValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isRoleValid {
		message := "Invalid email! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Create new user with random password
	password, errCode, err := controller.Service.NewUserWithEmail(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.NewUserWithEmailResponse{
			Email:    requestData.Email,
			Password: password,
			Role:     requestData.Role,
		},
	})
}

// @Tags Add new user - [super-admin]
// @Summary Add new user with phone number
// @Accept  json
// @Produce  json
// @Param   email path string true "Enter your phone number"
// @Param   role path string true "Select role" Enums(super-admin, admin, manager, manager-assist, deliver, customer, customer-service)
// @Success 200 {object} response.NewUserWithEmailResponse "OK"
// @Failure 400 {string} string "Invalid phone number or role!"
// @Failure 302 {string} string "User with this phone number already exists!"
// @Router /add-new-user-with-phone [post]
func (controller *AuthController) NewUserWithPhoneNumber(c *gin.Context) {
	// Get data of req body
	var requestData request.NewUserWithPhoneNumberRequest
	c.Bind(&requestData)
	isPhoneNumberValid := utils.IsPhoneNumberValid(requestData.PhoneNumber)
	isRoleValid := utils.IsRoleValid(requestData.Role)
	if !isPhoneNumberValid && !isRoleValid {
		message := "Invalid phone number and role! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isPhoneNumberValid {
		message := "Invalid phone number! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}
	if !isRoleValid {
		message := "Invalid phone role! Please enter valid information."
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("%s", message))
		return
	}

	// Create new user with random password
	password, errCode, err := controller.Service.NewUserWithPhoneNumber(&requestData)
	if err != nil {
		c.AbortWithError(errCode, err)
		return
	}

	// Return resp
	c.JSON(http.StatusOK, types.WebSuccessResponse{
		Data: response.NewUserWithPhoneNumberResponse{
			PhoneNumber: requestData.PhoneNumber,
			Password:    password,
			Role:        requestData.Role,
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
