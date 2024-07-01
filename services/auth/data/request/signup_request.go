package request

type SignUpEmailRequest struct {
	Email    string
	Password string
}

type SignUpPhoneNumberRequest struct {
	PhoneNumber int
	Password    string
}

type SignUpRequest struct {
	Email       string
	PhoneNumber int
	Password    string
}
