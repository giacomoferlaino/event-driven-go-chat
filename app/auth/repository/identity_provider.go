package repository

import (
	"chat/app/auth/config"
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type IdentityProvider interface {
	GetAccessToken(username string, password string) (string, error)
}

func NewKeycloak(basePath string, realm string) Keycloak {
	client := gocloak.NewClient(basePath)

	return Keycloak{
		client:        client,
		ctx:           context.Background(),
		realm:         realm,
		clientAccount: config.KcClient(),
	}
}

type Keycloak struct {
	client        *gocloak.GoCloak
	ctx           context.Context
	realm         string
	clientAccount gocloak.Client
}

func (k Keycloak) GetAccessToken(username string, password string) (string, error) {
	jwt, err := k.client.Login(k.ctx, *k.clientAccount.ClientID, *k.clientAccount.Secret, k.realm, username, password)
	if err != nil {
		return "", err
	}
	return jwt.AccessToken, nil
}
