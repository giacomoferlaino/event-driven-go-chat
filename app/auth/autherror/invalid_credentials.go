package autherror

import (
	"chat/pkg/api"
	"net/http"
)

func NewInvalidCredentials(rootError error) *api.HttpError {
	return api.NewHTTPError("invalid_credentials", http.StatusUnauthorized, rootError)
}
