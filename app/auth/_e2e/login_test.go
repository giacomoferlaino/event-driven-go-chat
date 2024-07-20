package e2e

import (
	"chat/app/auth/graph/generated/e2e"
	"chat/pkg/test"
	"errors"
	"testing"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestLogin(t *testing.T) {
	t.Run("if the credentials are valid", func(t *testing.T) {
		t.Run("it should return the JWT", func(t *testing.T) {
			input := e2e.UserCredentials{
				Username: *chatUser().Username,
				Password: *chatUser().Password,
			}
			res, err := e2e.Login(e2eEnv.GQLClient.Ctx, e2eEnv.GQLClient.Client, input)

			want := e2e.LoginLoginJWT{
				AccessToken: "",
			}
			got := res.GetLogin()
			test.AssertNotEqual(want, got, t)
			test.AssertEqual(nil, err, t)
		})
	})

	t.Run("if the credetntials are not valid", func(t *testing.T) {
		t.Run("it should return an error", func(t *testing.T) {
			input := e2e.UserCredentials{
				Username: "invalid_user",
				Password: "invalid_password",
			}
			res, err := e2e.Login(e2eEnv.GQLClient.Ctx, e2eEnv.GQLClient.Client, input)
			var gqlError *gqlerror.Error
			if !errors.As(err, &gqlError) {
				t.Errorf("The returned error type is not gqlerror.Error")
			}

			want := e2e.LoginLoginJWT{
				AccessToken: "",
			}
			got := res.GetLogin()
			test.AssertEqual(want, got, t)
			test.AssertEqual("invalid_credentials", gqlError.Message, t)
		})
	})
}
