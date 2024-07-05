package backend

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ErrorResponse struct{ Response }

func NewErrorResponse(msg string) *ErrorResponse {
	return &ErrorResponse{Response: Response{
		Message: msg,
		Status:  "error",
	}}
}

func NewResponse(msg string, status string) *Response {
	return &Response{
		Message: msg,
		Status:  status,
	}
}
