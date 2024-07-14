package types

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessPaginatedResponse struct {
	Filter     *Filter     `json:"filter"`
	Pagination *Pagination `json:"pagination"`
	Data       any         `json:"data"`
}
