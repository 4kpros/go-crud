package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/config"
	"github.com/4kpros/go-api/features/auth/data/request"
	"github.com/4kpros/go-api/features/auth/repository"
	"github.com/4kpros/go-api/features/user/model"
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
	if !types.AuthProviders[reqData.Provider] {
		message := "Invalid provider! Please enter valid information."
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", message)
		return
	}
	var providerUserId = "Test"
	if len(providerUserId) <= 0 {
		message := "Invalid provider or token! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}
	newUserFound, errFound := service.Repository.FindByProvider(reqData.Provider, providerUserId)
	if errFound != nil || newUserFound == nil || newUserFound.Provider != reqData.Provider {
		// Save new user
		user := &model.User{
			Provider:       reqData.Provider,
			ProviderUserId: providerUserId,
		}
		err = service.Repository.Create(user)
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

func (service *AuthServiceImpl) NewUserWithEmail(reqData *request.NewUserWithEmailRequest) (password string, errCode int, err error) {
	// Check if user exists
	user, errFound := service.Repository.FindByEmail(reqData.Email)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if user != nil && user.Email == reqData.Email {
		message := "User with this email already exists! Please use another email."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	var randomPassword = utils.GenerateRandomPassword(8)
	user.Email = reqData.Email
	user.Password = randomPassword
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	password = randomPassword

	return
}

func (service *AuthServiceImpl) NewUserWithPhoneNumber(reqData *request.NewUserWithPhoneNumberRequest) (password string, errCode int, err error) {
	// Check if user exists
	user, errFound := service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if user != nil && user.PhoneNumber == reqData.PhoneNumber {
		message := "User with this phone number already exists! Please use another phone number."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	var randomPassword = utils.GenerateRandomPassword(8)
	user.PhoneNumber = reqData.PhoneNumber
	user.Password = randomPassword
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	password = randomPassword

	return
}

func (service *AuthServiceImpl) SignUpWithEmail(reqData *request.SignUpWithEmailRequest) (token string, errCode int, err error) {
	// Check if user exists
	user, errFound := service.Repository.FindByEmail(reqData.Email)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if user != nil && user.Email == reqData.Email {
		message := "User with this email already exists! Please use another email."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	user.Email = reqData.Email
	user.Password = reqData.Password
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Generate new token
	var expires = utils.NewOthersExpiresDate()
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  user.ID,
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
	user, errFound := service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
	if errFound != nil {
		errCode = http.StatusInternalServerError
		err = errFound
		return
	}
	if user != nil && user.PhoneNumber == reqData.PhoneNumber {
		message := "User with this phone number already exists! Please use another phone number."
		errCode = http.StatusFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Create new user
	user.PhoneNumber = reqData.PhoneNumber
	user.Password = reqData.Password
	err = service.Repository.Create(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}

	// Generate new token
	var expires = utils.NewOthersExpiresDate()
	jwt, errJwt := utils.EncryptJWTToken(
		&types.JwtToken{
			UserId:  user.ID,
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
	// Extract token information
	jwt, errJwt := utils.DecryptJWTToken(reqData.Token, config.AppPem.JwtPublicKey)
	if errJwt != nil {
		errCode = http.StatusNotFound
		err = errJwt
		return
	}
	if jwt == nil || jwt.UserId <= 0 {
		message := "Invalid token or expired! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Check if user exists
	var userId = fmt.Sprintf("%d", jwt.UserId)
	user, errFound := service.Repository.FindById(userId)
	if errFound != nil || user == nil {
		message := "Invalid token! Please enter valid information."
		errCode = http.StatusNotFound
		err = fmt.Errorf("%s", message)
		return
	}

	// Check if account is activated
	if user.IsActivated {
		message := "User account is already activated! Please sign in and start using our services."
		errCode = http.StatusForbidden
		err = fmt.Errorf("%s", message)
		return
	}

	// Update account validation status
	tmpActivatedAt := time.Now()
	user.ActivatedAt = &tmpActivatedAt
	user.IsActivated = true
	err = service.Repository.Update(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	// Create user info
	var userInfo = &model.UserInfo{}
	err = service.Repository.CreateUserInfo(userInfo)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	// Update user with user info id
	user.UserInfoId = userInfo.ID
	err = service.Repository.Update(user)
	if err != nil {
		errCode = http.StatusInternalServerError
		return
	}
	activatedAt = user.ActivatedAt

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
	var user *model.User
	var message string
	if len(reqData.Email) > 0 {
		user, err = service.Repository.FindByEmail(reqData.Email)
		message = "No user found with this email! Please enter valid information."
	} else {
		user, err = service.Repository.FindByPhoneNumber(reqData.PhoneNumber)
		message = "No user found with this phone number! Please enter valid information."
	}
	if err != nil || user == nil || (user.Email != reqData.Email && user.PhoneNumber != reqData.PhoneNumber) {
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
	var user model.User

	// Update password
	user.ID = userId
	user.Password = reqData.NewPassword
	service.Repository.Update(&user)

	return
}
