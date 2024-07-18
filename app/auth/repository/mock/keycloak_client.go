package mock

import (
	"chat/pkg/test"
	"context"
	"errors"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloackClientMock() KeycloackClientMock {
	return KeycloackClientMock{
		LoginReturn: &test.ReturnTuple[*gocloak.JWT, error]{
			Val1: nil,
			Val2: errors.New(""),
		},
	}
}

type KeycloackClientMock struct {
	LoginReturn *test.ReturnTuple[*gocloak.JWT, error]
}

func (k KeycloackClientMock) Login(ctx context.Context, clientID, clientSecret, realm, username, password string) (*gocloak.JWT, error) {
	return k.LoginReturn.Val1, k.LoginReturn.Val2
}
