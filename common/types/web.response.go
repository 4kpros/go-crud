package types

type WebErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type WebSuccessResponse struct {
	Data any `json:"data"`
}

type WebSuccessPaginatedResponse struct {
	Data       any        `json:"data"`
	Filter     Filter     `json:"filter"`
	Pagination Pagination `json:"pagination"`
}
