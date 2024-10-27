package httperror

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%d error: %s", e.StatusCode, e.Message)
}

func InternalServerError() *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    "unexpected error occurred",
	}
}

func BadRequestError(details any) *HttpError {
	return &HttpError{
		StatusCode: http.StatusBadRequest,
		Message:    "bad request",
		Details:    details,
	}
}

func ConflictError(details any) *HttpError {
	return &HttpError{
		StatusCode: http.StatusConflict,
		Message:    "conflict with existing data",
		Details:    details,
	}
}

func NotFoundError(details any) *HttpError {
	return &HttpError{
		StatusCode: http.StatusNotFound,
		Message:    "data not found",
		Details:    details,
	}
}
