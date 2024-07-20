package repository

import (
	"chat/app/auth/autherror"
	"chat/app/auth/config"
	"context"
	"errors"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
)

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
	client        KeycloakClient
	ctx           context.Context
	realm         string
	clientAccount gocloak.Client
}

func (k Keycloak) GetAccessToken(username string, password string) (string, error) {
	jwt, err := k.client.Login(k.ctx, *k.clientAccount.ClientID, *k.clientAccount.Secret, k.realm, username, password)

	if err != nil {
		var gkError *gocloak.APIError
		if errors.As(err, &gkError) {
			if gkError.Code == http.StatusUnauthorized {
				return "", autherror.NewInvalidCredentials(*gkError)
			}
		}

		return "", err
	}

	return jwt.AccessToken, nil
}
