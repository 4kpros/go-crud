package service

import (
	"time"

	"github.com/4kpros/go-api/features/auth/data/request"
)

type AuthService interface {
	SignIn(deviceName string, reqData *request.SignInRequest) (validateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error)
	SignInWithProvider(deviceName string, reqData *request.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error)

	SignUp(reqData *request.SignUpRequest) (token string, errCode int, err error)

	ActivateAccount(reqData *request.ActivateAccountRequest) (date *time.Time, errCode int, err error)

	ResetPasswordInit(reqData *request.ResetPasswordInitRequest) (token string, errCode int, err error)
	ResetPasswordCode(reqData *request.ResetPasswordCodeRequest) (token string, errCode int, err error)
	ResetPasswordNewPassword(reqData *request.ResetPasswordNewPasswordRequest) (errCode int, err error)
}
