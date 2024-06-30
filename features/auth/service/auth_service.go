package service

import (
	"time"

	"github.com/4kpros/go-api/features/auth/data/request"
)

type AuthService interface {
	SignInWithEmail(reqData *request.SignInWithEmailRequest) (validateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error)
	SignInWithPhoneNumber(reqData *request.SignInWithPhoneNumberRequest) (validateAccountToken string, accessToken string, accessExpires *time.Time, errCode int, err error)
	SignInWithProvider(reqData *request.SignInWithProviderRequest) (accessToken string, accessExpires *time.Time, errCode int, err error)

	NewUserWithEmail(reqData *request.NewUserWithEmailRequest) (password string, errCode int, err error)
	NewUserWithPhoneNumber(reqData *request.NewUserWithPhoneNumberRequest) (password string, errCode int, err error)

	SignUpWithEmail(reqData *request.SignUpWithEmailRequest) (token string, errCode int, err error)
	SignUpWithPhoneNumber(reqData *request.SignUpWithPhoneNumberRequest) (token string, errCode int, err error)

	ActivateAccount(reqData *request.ActivateAccountRequest) (date *time.Time, errCode int, err error)

	ResetPasswordInit(reqData *request.ResetPasswordInitRequest) (token string, errCode int, err error)
	ResetPasswordCode(reqData *request.ResetPasswordCodeRequest) (token string, errCode int, err error)
	ResetPasswordNewPassword(reqData *request.ResetPasswordNewPasswordRequest) (errCode int, err error)
}
