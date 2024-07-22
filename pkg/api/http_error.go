package api

import (
	"fmt"
	"net/http"
)

func NewHTTPError(message string, statusCode int, err error) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

type HttpError struct {
	Message    string
	Err        error
	StatusCode int
}

func (h HttpError) Error() string {
	return fmt.Errorf("http_error %d: %w", h.StatusCode, h.Err).Error()
}

func NewInternalServerError(err error) HttpError {
	return HttpError{
		Message:    "internal_server_error",
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}
