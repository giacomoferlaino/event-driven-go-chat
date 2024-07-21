package mock

import (
	"chat/app/auth/domain"
	"chat/pkg/test"
	"errors"
)

func NewIdentityProvider() IdentityProvider {
	return IdentityProvider{
		GetAccessTokenReturn: &test.ReturnTuple[domain.JWT, error]{
			Val1: domain.JWT{},
			Val2: errors.New(""),
		},
	}
}

type IdentityProvider struct {
	GetAccessTokenReturn *test.ReturnTuple[domain.JWT, error]
}

func (i IdentityProvider) GetJWT(username string, password string) (domain.JWT, error) {
	return i.GetAccessTokenReturn.Val1, i.GetAccessTokenReturn.Val2
}
