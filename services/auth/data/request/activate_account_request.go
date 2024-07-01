package request

type ActivateAccountRequest struct {
	Token string
	Code  int
}
