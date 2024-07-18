package repository

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakClient interface {
	Login(ctx context.Context, clientID, clientSecret, realm, username, password string) (*gocloak.JWT, error)
}
