package response

type AddNewUserWithEmailResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AddNewUserWithPhoneNumberResponse struct {
	PhoneNumber int    `json:"phoneNumber"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
