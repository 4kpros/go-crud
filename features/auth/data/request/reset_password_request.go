package request

type ResetPasswordEmailInitRequest struct {
	Email string
}
type ResetPasswordPhoneNumberInitRequest struct {
	PhoneNumber int
}

type ResetPasswordInitRequest struct {
	Email       string
	PhoneNumber int
}

type ResetPasswordCodeRequest struct {
	Token string
	Code  int
}

type ResetPasswordNewPasswordRequest struct {
	Token       string
	NewPassword string
}
