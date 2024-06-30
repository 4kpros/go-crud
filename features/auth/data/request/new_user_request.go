package request

type NewUserWithEmailRequest struct {
	Email string
	Role  string
}

type NewUserWithPhoneNumberRequest struct {
	PhoneNumber int
	Role        string
}
