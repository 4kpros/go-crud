package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/auth/data/request"
	modelNewUser "github.com/4kpros/go-api/features/auth/model"
	"github.com/4kpros/go-api/features/auth/repository"
	modelUser "github.com/4kpros/go-api/features/user/model"
)

type AuthServiceImpl struct {
	Repository repository.AuthRepository
}

func NewAuthServiceImpl(repository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{Repository: repository}
}

func (service *AuthServiceImpl) SignInWithEmail(reqData *request.SignInWithEmailRequest) (activateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Check if user exists
	newUserFound, errFound := service.Repository.FindByEmail(reqData.Email)
	isPasswordMatches, errCompare := utils.CompareToArgon2id(reqData.Password, newUserFound.Password)
	if errFound != nil || errCompare != nil || newUserFound == nil || newUserFound.Email != reqData.Email || !isPasswordMatches {
		message := "Invalid email address or password! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Check if account is activated
	if !newUserFound.IsActivated {
		message := "Account found but not activated! Please activate your account before."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", message)
		var expires = utils.NewOthersExpiresDate()
		jwt, errJwt := utils.EncryptJWTToken(
			&types.JwtToken{
				UserId:  newUserFound.ID,
				Expires: *expires,
				Device:  "NA",
			},
			config.AppPem.JwtPrivateKey,
		)
		if errJwt != nil {
			errCode = http.StatusInternalServerError
			err = errJwt
		}
		activateAccountToken = jwt
		return
	}

	// Generate new token
	var expires = utils.NewExpiresDate(reqData.StayConnected)
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  newUserFound.ID,
			Expires: *expires,
			Device:  "NA",
		},
		config.AppPem.JwtPrivateKey,
	)
	if errJwt != nil {
		errCode = http.StatusInternalServerError
		err = errJwt
	}
	accessToken = jwt
	accessExpires = expires

	return
}

func (service *AuthServiceImpl) SignInWithPhoneNumber(reqData *request.SignInWithPhoneNumberRequest) (activateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error) {
	// Check if user exists
	newUserFound, errFound := service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	isPasswordMatches, errCompare := utils.CompareToArgon2id(reqData.Password, newUserFound.Password)
	if errFound != nil || errCompare != nil || newUserFound == nil || newUserFound.PhoneNumber != reqData.PhoneNumber || !isPasswordMatches {
		message := "Invalid phone number or password! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Check if account is activated
	if !newUserFound.IsActivated {
		message := "Account found but not activated! Please activate your account before."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", message)
		var expires = utils.NewOthersExpiresDate()
		jwt, errJwt := utils.EncryptJWTToken(
			&types.JwtToken{
				UserId:  newUserFound.ID,
				Expires: *expires,
				Device:  "NA",
			},
			config.AppPem.JwtPrivateKey,
		)
		if errJwt != nil {
			errCode = http.StatusInternalServerError
			err = errJwt
		}
		activateAccountToken = jwt
		return
	}

	// Generate new token
	var expires = utils.NewExpiresDate(reqData.StayConnected)
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  newUserFound.ID,
			Expires: *expires,
			Device:  "NA",
		},
		config.AppPem.JwtPrivateKey)
	if errJwt != nil {
		errCode = http.StatusInternalServerError
		err = errJwt
	}
	accessToken = jwt
	accessExpires = expires

	return
}

func (service *AuthServiceImpl) SignInWithProvider(reqData *request.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error) {
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
		newUser := &modelNewUser.NewUser{
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
	var expires = utils.NewExpiresDate(true)
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  newUserFound.ID,
			Expires: *expires,
			Device:  "NA",
		},
		config.AppPem.JwtPrivateKey,
	)
	if errJwt != nil {
		errCode = http.StatusInternalServerError
		err = errJwt
	}
	accessToken = jwt
	accessExpires = expires

	return
}

func (service *AuthServiceImpl) AddNewUserWithEmail(reqData *request.AddNewUserWithEmailRequest) (password string, errCode int, err error) {
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
	var randomPassword = utils.GenerateRandomPassword(8)
	newUser.Email = reqData.Email
	newUser.Password = randomPassword
	err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	password = randomPassword

	return
}

func (service *AuthServiceImpl) AddNewUserWithPhoneNumber(reqData *request.AddNewUserWithPhoneNumberRequest) (password string, errCode int, err error) {
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
	var randomPassword = utils.GenerateRandomPassword(8)
	newUser.PhoneNumber = reqData.PhoneNumber
	newUser.Password = randomPassword
	err = service.Repository.Create(newUser)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	password = randomPassword

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
	var expires = utils.NewOthersExpiresDate()
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  newUser.ID,
			Expires: *expires,
			Device:  "NA",
		},
		config.AppPem.JwtPrivateKey,
	)
	if errJwt != nil {
		errCode = http.StatusInternalServerError
		err = errJwt
	}
	token = jwt

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
	var expires = utils.NewOthersExpiresDate()
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  newUser.ID,
			Expires: *expires,
			Device:  "NA",
		},
		config.AppPem.JwtPrivateKey,
	)
	if errJwt != nil {
		errCode = http.StatusInternalServerError
		err = errJwt
	}
	token = jwt

	return
}

func (service *AuthServiceImpl) ActivateAccount(reqData *request.ActivateAccountRequest) (activatedAt *time.Time, errCode int, err error) {
	// Check if the token is valid and extract userId
	jwt, errJwt := utils.DecryptJWTToken(reqData.Token, config.AppPem.JwtPublicKey)
	if errJwt != nil || jwt == nil {
		message := "Invalid token! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}
	var userId = fmt.Sprintf("%d", jwt.UserId)
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
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	// Create user account details
	var user = &modelUser.User{}
	user.NewUserId = newUser.ID
	err = service.Repository.CreateUserAccountDetails(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
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
	var newUser *modelNewUser.NewUser
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
	var newUser modelNewUser.NewUser

	// Update password
	newUser.ID = userId
	newUser.Password = reqData.NewPassword
	service.Repository.Update(&newUser)

	return
}
