package e2e

import (
	"chat/app/auth/graph/generated/e2e"
	"chat/pkg/test"
	"testing"
)

func TestLogin(t *testing.T) {
	input := e2e.UserCredentials{
		Username: "example@email.co",
		Password: "password",
	}
	res, err := e2e.Login(testEnv.GQLClient.Ctx, testEnv.GQLClient.Client, input)
	if err != nil {
		t.Fatal(err)
	}

	want := e2e.LoginLoginUser{
		FirstName: "James",
		LastName:  "Bond",
	}
	got := res.GetLogin()
	test.AssertEqual(want, got, t)
}
