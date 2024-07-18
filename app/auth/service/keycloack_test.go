package service

import (
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
			mockedReturn := test.ReturnTuple[string, error]{
				Val1: "",
				Val2: errors.New("Error retrieving the token"),
			}
			defer test.Stub(identityProvider.GetAccessTokenReturn, mockedReturn)()

			accessToken, err := keycloakService.Login("username", "password")

			test.AssertEqual(mockedReturn.Val1, accessToken, t)
			test.AssertEqual(mockedReturn.Val2, err, t)
		})
	})

	t.Run("if the GetAccessToken succeeds", func(t *testing.T) {
		t.Run("it should return the access token", func(t *testing.T) {
			mockedReturn := test.ReturnTuple[string, error]{Val1: "access_token", Val2: nil}
			defer test.Stub(identityProvider.GetAccessTokenReturn, mockedReturn)()

			accessToken, err := keycloakService.Login("username", "password")

			test.AssertEqual(mockedReturn.Val1, accessToken, t)
			test.AssertEqual(mockedReturn.Val2, err, t)
		})
	})
}
