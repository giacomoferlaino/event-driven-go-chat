package e2e

import (
	"chat/app/auth/graph/generated/e2e"
	"chat/pkg/test"
	"testing"
)

func TestLogin(t *testing.T) {
	input := e2e.UserCredentials{
		Username: *chatUser().KCUser.Username,
		Password: *chatUser().Password,
	}
	res, err := e2e.Login(e2eEnv.GQLClient.Ctx, e2eEnv.GQLClient.Client, input)
	if err != nil {
		t.Fatal(err)
	}

	want := e2e.LoginLoginJWT{
		AccessToken: "",
	}
	got := res.GetLogin()
	test.AssertNotEqual(want, got, t)
}
