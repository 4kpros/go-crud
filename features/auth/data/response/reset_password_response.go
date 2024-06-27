package response

type ResetPasswordInitResponse struct {
	Token string
}

type ResetPasswordCodeResponse struct {
	Token string
}

type ResetPasswordNewPasswordResponse struct {
	Message string
}
