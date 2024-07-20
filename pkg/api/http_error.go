package api

import "net/http"

func NewHTTPError(message string, statusCode uint, rootError error) HttpError {
	return HttpError{
		Message:    message,
		StatusCode: statusCode,
		RootError:  rootError,
	}
}

type HttpError struct {
	Message    string
	RootError  error
	StatusCode uint
}

func (h HttpError) Error() string {
	return h.Message
}

func NewInternalServerError(rootError error) HttpError {
	return HttpError{
		Message:    "internal_server_error",
		StatusCode: http.StatusInternalServerError,
		RootError:  rootError,
	}
}
