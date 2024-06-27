package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/features/auth/data/request"
	"github.com/4kpros/go-api/features/auth/model"
	"github.com/4kpros/go-api/features/auth/repository"
)

type AuthServiceImpl struct {
	Repository repository.AuthRepository
}

func NewAuthServiceImpl(repository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{Repository: repository}
}

func (service *AuthServiceImpl) SignInWithEmail(reqData *request.SignInWithEmailRequest) (activateAccountToken string, accessToken string, refreshToken string, accessExpires string, refreshExpires string, errCode int, err error) {
	// Check if user exists
	newUserFound, errFound := service.Repository.FindByEmail(reqData.Email)
	isPasswordMatches, errCompare := utils.CompareToArgon2id(reqData.Password, newUserFound.Password)
	if errFound != nil || errCompare != nil || newUserFound == nil || newUserFound.Email != reqData.Email || !isPasswordMatches {
		message := "Invalid email address or password! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Check if user account is activated
	if !newUserFound.IsActivated {
		message := "Account found but not activated! Please activate your account before."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", message)
		activateAccountToken = "EOF"
		return
	}

	// Generate new token
	accessToken = "EOF"
	refreshToken = "EOF"
	accessExpires = "EOF"
	refreshExpires = "EOF"

	return
}

func (service *AuthServiceImpl) SignInWithPhoneNumber(reqData *request.SignInWithPhoneNumberRequest) (activateAccountToken string, accessToken string, refreshToken string, accessExpires string, refreshExpires string, errCode int, err error) {
	// Check if user exists
	newUserFound, errFound := service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	isPasswordMatches, errCompare := utils.CompareToArgon2id(reqData.Password, newUserFound.Password)
	if errFound != nil || errCompare != nil || newUserFound == nil || newUserFound.PhoneNumber != reqData.PhoneNumber || !isPasswordMatches {
		message := "Invalid phone number or password! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Check if user account is activated
	if !newUserFound.IsActivated {
		message := "Account found but not activated! Please activate your account before."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", message)
		activateAccountToken = "EOF"
		return
	}

	// Generate new token
	accessToken = "EOF"
	refreshToken = "EOF"
	accessExpires = "EOF"
	refreshExpires = "EOF"

	return
}

func (service *AuthServiceImpl) SignInWithProvider(reqData *request.SignInWithProviderRequest) (accessToken string, refreshToken string, accessExpires string, refreshExpires string, errCode int, err error) {
	// Validate provider token and update user
	var providerUserId string
	switch reqData.Provider {
	case "google":
		providerUserId = "googleUserId"
	case "facebook":
		providerUserId = "facebookUserId"
	}
	if len(providerUserId) <= 0 {
		message := "Invalid provider or token! Please use valid provider ['google', 'facebook']."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}
	newUserFound, errFound := service.Repository.FindByProvider(reqData.Provider, providerUserId)
	if errFound != nil || newUserFound == nil || newUserFound.Provider != reqData.Provider {
		// Save new user
		newUser := &model.NewUser{
			Provider:       reqData.Provider,
			ProviderUserId: providerUserId,
		}
		err = service.Repository.Create(newUser)
		if err != nil {
			errCode = http.StatusInternalServerError
			return
		}
	}

	// Generate new token
	accessToken = "EOF"
	refreshToken = "EOF"
	accessExpires = "EOF"
	refreshExpires = "EOF"

	return
}

func (service *AuthServiceImpl) SignUpWithEmail(reqData *request.SignUpWithEmailRequest) (token string, errCode int, err error) {
	// Check if user exists
	newUser, errFound := service.Repository.FindByEmail(reqData.Email)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if newUser != nil && newUser.Email == reqData.Email {
		message := "User with this email already exists! Please use another email."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	newUser.Email = reqData.Email
	newUser.Password = reqData.Password
	err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Generate new token
	token = "EOF"

	return
}

func (service *AuthServiceImpl) SignUpWithPhoneNumber(reqData *request.SignUpWithPhoneNumberRequest) (token string, errCode int, err error) {
	// Check if user exists
	newUser, errFound := service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if newUser != nil && newUser.PhoneNumber == reqData.PhoneNumber {
		message := "User with this phone number already exists! Please use another phone number."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	newUser.PhoneNumber = reqData.PhoneNumber
	newUser.Password = reqData.Password
	err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Generate new token
	token = "EOF"

	return
}

func (service *AuthServiceImpl) ActivateAccount(reqData *request.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Check if the token is valid and extract userId
	var userId string
	newUser, errFound := service.Repository.FindById(userId)
	if errFound != nil || newUser == nil {
		message := "Invalid token! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Update account validation status
	newUser.ActivatedAt = &time.Time{}
	err = service.Repository.Update(newUser)
	activatedAt = newUser.ActivatedAt

	return
}

func (service *AuthServiceImpl) ResetPasswordCode(reqData *request.ResetPasswordCodeRequest) (token string, errCode int, err error) {
	// Check if the token is valid and extract code + userId

	// Check if code is valid

	// Generate new token
	token = "EOF"

	return
}

func (service *AuthServiceImpl) ResetPasswordInit(reqData *request.ResetPasswordInitRequest) (token string, errCode int, err error) {
	// Check if user exists
	var newUser *model.NewUser
	var message string
	if len(reqData.Email) > 0 {
		newUser, err = service.Repository.FindByEmail(reqData.Email)
		message = "No user found with this email! Please enter valid information."
	} else {
		newUser, err = service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
		message = "No user found with this phone number! Please enter valid information."
	}
	if err != nil || newUser == nil || (newUser.Email != reqData.Email && newUser.PhoneNumber != reqData.PhoneNumber) {
		err = fmt.Errorf("%s", message)
		return
	}

	// Generate new token
	token = "EOF"

	return
}

func (service *AuthServiceImpl) ResetPasswordNewPassword(reqData *request.ResetPasswordNewPasswordRequest) (errCode int, err error) {
	// Check if the token is valid
	var userId uint
	var newUser model.NewUser

	// Update password
	newUser.ID = userId
	newUser.Password = reqData.NewPassword
	service.Repository.Update(&newUser)

	return
}
