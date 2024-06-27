package request

type ResetPasswordInitRequest struct {
	Email       string
	PhoneNumber int
}

type ResetPasswordCodeRequest struct {
	Token string
}

type ResetPasswordNewPasswordRequest struct {
	Token       string
	NewPassword string
}
