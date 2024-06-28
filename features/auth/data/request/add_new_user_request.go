package request

type AddNewUserWithEmailRequest struct {
	Email string
	Role  string
}

type AddNewUserWithPhoneNumberRequest struct {
	PhoneNumber int
	Role        string
}
