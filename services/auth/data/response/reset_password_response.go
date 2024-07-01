package response

type ResetPasswordInitResponse struct {
	Token string `json:"token"`
}

type ResetPasswordCodeResponse struct {
	Token string `json:"token"`
}

type ResetPasswordNewPasswordResponse struct {
	Message string `json:"message"`
}
