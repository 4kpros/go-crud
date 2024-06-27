package request

type SignUpWithEmailRequest struct {
	Email    string
	Password string
}

type SignUpWithPhoneNumberRequest struct {
	PhoneNumber int
	Password    string
}

type SignUpWithProviderRequest struct {
	Provider string
	Token    string
}
