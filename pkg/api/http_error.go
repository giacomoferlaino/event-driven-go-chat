package api

import (
	"fmt"
	"net/http"
)

func NewHTTPError(message string, statusCode int, rootError error) *HttpError {
	return &HttpError{
		Message:    message,
		StatusCode: statusCode,
		RootError:  rootError,
	}
}

type HttpError struct {
	Message    string
	RootError  error
	StatusCode int
}

func (h HttpError) Error() string {
	return fmt.Errorf("http_error %d: %w", h.StatusCode, h.RootError).Error()
}

func NewInternalServerError(rootError error) HttpError {
	return HttpError{
		Message:    "internal_server_error",
		StatusCode: http.StatusInternalServerError,
		RootError:  rootError,
	}
}
