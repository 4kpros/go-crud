package response

type NewUserWithEmailResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type NewUserWithPhoneNumberResponse struct {
	PhoneNumber int    `json:"phoneNumber"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
