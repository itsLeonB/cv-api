package dto

type JsonResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Err     string `json:"error,omitempty"`
}

func NewSuccessResponse(data any) *JsonResponse {
	return &JsonResponse{
		Success: true,
		Data:    data,
	}
}
