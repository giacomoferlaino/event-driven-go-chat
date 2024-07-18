package mock

import (
	"chat/pkg/test"
	"errors"
)

func NewIdentityProvider() IdentityProvider {
	return IdentityProvider{
		GetAccessTokenReturn: &test.ReturnTuple[string, error]{
			Val1: "",
			Val2: errors.New(""),
		},
	}
}

type IdentityProvider struct {
	GetAccessTokenReturn *test.ReturnTuple[string, error]
}

func (i IdentityProvider) GetAccessToken(username string, password string) (string, error) {
	return i.GetAccessTokenReturn.Val1, i.GetAccessTokenReturn.Val2
}
