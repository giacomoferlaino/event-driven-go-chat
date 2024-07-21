package service

import (
	"chat/app/auth/domain"
	"chat/app/auth/service/mock"
	"chat/pkg/test"
	"errors"
	"testing"
)

func TestLogin(t *testing.T) {
	identityProvider := mock.NewIdentityProvider()
	keycloakService := NewKeycloak(identityProvider)

	t.Run("if the GetAccessToken fails", func(t *testing.T) {
		t.Run("it should forward the identity provider error", func(t *testing.T) {
			mockedReturn := test.ReturnTuple[domain.JWT, error]{
				Val1: domain.JWT{},
				Val2: errors.New("Error retrieving the token"),
			}
			defer test.Stub(identityProvider.GetAccessTokenReturn, mockedReturn)()

			jwt, err := keycloakService.Login("username", "password")

			test.AssertEqual(mockedReturn.Val1, jwt, t)
			test.AssertEqual(mockedReturn.Val2, err, t)
		})
	})

	t.Run("if the GetAccessToken succeeds", func(t *testing.T) {
		t.Run("it should return the access token", func(t *testing.T) {
			mockedReturn := test.ReturnTuple[domain.JWT, error]{
				Val1: domain.JWT{
					AccessToken:  "access_token",
					RefreshToken: "refresh_token",
				},
				Val2: nil,
			}
			defer test.Stub(identityProvider.GetAccessTokenReturn, mockedReturn)()

			jwt, err := keycloakService.Login("username", "password")

			test.AssertEqual(mockedReturn.Val1, jwt, t)
			test.AssertEqual(mockedReturn.Val2, err, t)
		})
	})
}
