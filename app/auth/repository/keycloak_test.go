package repository

import (
	"chat/app/auth/config"
	"chat/app/auth/domain"
	"chat/app/auth/repository/mock"
	"chat/pkg/test"
	"context"
	"errors"
	"testing"

	"github.com/Nerzal/gocloak/v13"
)

func TestGetJWT(t *testing.T) {
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

			jwt, err := keycloakRepository.GetJWT("username", "password")

			test.AssertEqual(domain.JWT{}, jwt, t)
			test.AssertEqual(mockedReturn.Val2, err, t)
		})
	})

	t.Run("if the login return teh JWT", func(t *testing.T) {
		t.Run("it should return the JWT access token", func(t *testing.T) {
			mockedReturn := test.ReturnTuple[*gocloak.JWT, error]{
				Val1: &gocloak.JWT{AccessToken: "access_token", RefreshToken: "refresh_token"},
				Val2: nil,
			}
			defer test.Stub(keycloakClient.LoginReturn, mockedReturn)()

			jwt, err := keycloakRepository.GetJWT("username", "password")

			want := domain.JWT{
				AccessToken:  mockedReturn.Val1.AccessToken,
				RefreshToken: mockedReturn.Val1.RefreshToken,
			}
			test.AssertEqual(want, jwt, t)
			test.AssertEqual(nil, err, t)
		})
	})
}
