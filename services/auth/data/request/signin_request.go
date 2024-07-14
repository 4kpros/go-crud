package request

type SignInEmailRequest struct {
	Email         string
	Password      string
	StayConnected bool
}
type SignInPhoneNumberRequest struct {
	PhoneNumber   int
	Password      string
	StayConnected bool
}

type SignInRequest struct {
	Email         string
	PhoneNumber   int
	Password      string
	StayConnected bool
}

type SignInWithProviderRequest struct {
	Provider string
	Token    string
}
