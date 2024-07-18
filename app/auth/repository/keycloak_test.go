package repository

import (
	"chat/app/auth/config"
	"chat/app/auth/repository/mock"
	"chat/pkg/test"
	"context"
	"errors"
	"testing"

	"github.com/Nerzal/gocloak/v13"
)

func TestGetAccessToken(t *testing.T) {
	keycloakClient := mock.NewKeycloackClientMock()
	keycloakRepository := Keycloak{
		client:        keycloakClient,
		ctx:           context.Background(),
		realm:         "",
		clientAccount: config.KcClient(),
	}

	t.Run("if the login returns an error", func(t *testing.T) {
		t.Run("it should return an error", func(t *testing.T) {
			mockedReturn := test.ReturnTuple[*gocloak.JWT, error]{
				Val1: nil,
				Val2: errors.New("Error during login"),
			}
			defer test.Stub(keycloakClient.LoginReturn, mockedReturn)()

			accessToken, err := keycloakRepository.GetAccessToken("username", "password")

			test.AssertEqual("", accessToken, t)
			test.AssertEqual(mockedReturn.Val2, err, t)
		})
	})

	t.Run("if the login return teh JWT", func(t *testing.T) {
		t.Run("it should return the JWT access token", func(t *testing.T) {
			mockedReturn := test.ReturnTuple[*gocloak.JWT, error]{
				Val1: &gocloak.JWT{AccessToken: "access_token"},
				Val2: nil,
			}
			defer test.Stub(keycloakClient.LoginReturn, mockedReturn)()

			accessToken, err := keycloakRepository.GetAccessToken("username", "password")

			test.AssertEqual(mockedReturn.Val1.AccessToken, accessToken, t)
			test.AssertEqual(nil, err, t)
		})
	})
}
