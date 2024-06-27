package request

type SignInWithEmailRequest struct {
	Email    string
	Password string
}

type SignInWithPhoneNumberRequest struct {
	PhoneNumber int
	Password    string
}

type SignInWithProviderRequest struct {
	Provider string
	Token    string
}
