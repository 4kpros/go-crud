package response

type CreateWithEmailResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateWithPhoneNumberResponse struct {
	PhoneNumber int    `json:"phoneNumber"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
