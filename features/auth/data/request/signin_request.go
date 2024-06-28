package request

type SignInWithEmailRequest struct {
	Email         string
	Password      string
	StayConnected bool
}

type SignInWithPhoneNumberRequest struct {
	PhoneNumber   int
	Password      string
	StayConnected bool
}

type SignInWithProviderRequest struct {
	Provider string
	Token    string
}
