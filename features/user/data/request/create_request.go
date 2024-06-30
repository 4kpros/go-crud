package request

type CreateWithEmailRequest struct {
	Email string
	Role  string
}

type CreateWithPhoneNumberRequest struct {
	PhoneNumber int
	Role        string
}
