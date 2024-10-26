package model

type WebResponse struct {
	Success bool  `json:"success"`
	Data    any   `json:"data,omitempty"`
	Err     error `json:"error,omitempty"`
}

func NewSuccessResponse(data any) *WebResponse {
	return &WebResponse{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(err error) *WebResponse {
	return &WebResponse{
		Success: false,
		Err:     err,
	}
}
