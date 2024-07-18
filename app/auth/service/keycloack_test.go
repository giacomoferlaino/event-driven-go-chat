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
			want := errors.New("Error retrieving the token")
			mockedReturn := mock.GetAccessTokenReturn{Val: "", Err: want}
			defer test.Stub(identityProvider.GetAccessTokenReturn, mockedReturn)()

			_, got := keycloakService.Login("username", "password")

			test.AssertEqual(want, got, t)
		})
	})

	t.Run("if the GetAccessToken succeeds", func(t *testing.T) {
		t.Run("it should return the access token", func(t *testing.T) {
			want := "access_token"
			mockedReturn := mock.GetAccessTokenReturn{Val: want, Err: nil}
			defer test.Stub(identityProvider.GetAccessTokenReturn, mockedReturn)()

			got, _ := keycloakService.Login("username", "password")

			test.AssertEqual(want, got, t)

		})
	})
}
