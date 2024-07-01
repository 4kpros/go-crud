package request

type CreateWithEmailRequest struct {
	Email string
	Role  string
}

// Enums(super-admin, admin, manager, manager-assist, deliver, customer, customer-service)

type CreateWithPhoneNumberRequest struct {
	PhoneNumber int
	Role        string
}
