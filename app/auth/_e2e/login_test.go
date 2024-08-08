package e2e

import (
	"chat/app/auth/graph/generated/e2e"
	"chat/pkg/test"
	"chat/pkg/test/e2esdk"

	"errors"
	"testing"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestLogin(t *testing.T) {
	t.Run("if the credentials are valid", func(t *testing.T) {
		t.Run("it should return the JWT", func(t *testing.T) {
			input := e2e.UserCredentials{
				Username: *e2esdk.ChatUser().Username,
				Password: *e2esdk.ChatUser().Password,
			}
			res, err := e2e.Login(e2eEnv.GQLClient.Ctx, e2eEnv.GQLClient.Client, input)

			got := res.GetLogin()
			want := e2e.LoginLoginJWT{}
			test.AssertNotDeepEqual(want, got, t)
			test.AssertEqual(nil, err, t)
		})
	})

	t.Run("if the credentials are not valid", func(t *testing.T) {
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

			want := e2e.LoginLoginJWT{}
			got := res.GetLogin()
			test.AssertDeepEqual(want, got, t)
			test.AssertEqual("invalid_credentials", gqlError.Message, t)
		})
	})
}
