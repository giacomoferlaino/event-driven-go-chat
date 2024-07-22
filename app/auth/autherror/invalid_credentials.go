package autherror

import (
	"chat/pkg/api"
	"net/http"
)

func NewInvalidCredentials(err error) *api.HttpError {
	return api.NewHTTPError("invalid_credentials", http.StatusUnauthorized, err)
}
