package types

type WebErrorResponse struct {
	Message string `json:"message"`
}

type WebSuccessPaginatedResponse struct {
	Filter     *Filter     `json:"filter"`
	Pagination *Pagination `json:"pagination"`
	Data       any         `json:"data"`
}
