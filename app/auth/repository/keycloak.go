package repository

import (
	"chat/app/auth/autherror"
	"chat/app/auth/config"
	"chat/app/auth/domain"
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

func (k Keycloak) GetJWT(username string, password string) (domain.JWT, error) {
	gkJWT, err := k.client.Login(k.ctx, *k.clientAccount.ClientID, *k.clientAccount.Secret, k.realm, username, password)

	if err != nil {
		var gkError *gocloak.APIError
		if errors.As(err, &gkError) {
			if gkError.Code == http.StatusUnauthorized {
				return domain.JWT{}, autherror.NewInvalidCredentials(*gkError)
			}
		}

		return domain.JWT{}, err
	}

	jwt := domain.JWT{
		AccessToken:  gkJWT.AccessToken,
		RefreshToken: gkJWT.RefreshToken,
	}
	return jwt, nil
}
