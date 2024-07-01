package response

type SignUpResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
