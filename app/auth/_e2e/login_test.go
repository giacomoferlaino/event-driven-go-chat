package e2e

import (
	"chat/app/auth/graph/generated/e2e"
	"chat/app/auth/httpserver"
	"chat/pkg/env"
	"chat/pkg/keycloak/keycloacktest"
	"chat/pkg/test"
	"flag"
	"log"
	"net/http/httptest"
	"testing"
)

var testServer *httptest.Server
var gqlClient test.GQLClient

func TestMain(m *testing.M) {
	flag.Parse()
	env.Init()
	router := httpserver.Router()
	testServer = httptest.NewServer(router)
	gqlClient = test.NewGQLClient(testServer.URL)
	defer testServer.Close()

	kcTest, err := keycloacktest.New()
	if err != nil {
		log.Fatalln(err)
	}

	err = kcTest.Setup()
	defer kcTest.Teardown()
	if err != nil {
		log.Panicln(err)
	}
	m.Run()
}

func TestLogin(t *testing.T) {
	input := e2e.UserCredentials{
		Username: "example@email.co",
		Password: "password",
	}
	res, err := e2e.Login(gqlClient.Ctx, gqlClient.Client, input)
	if err != nil {
		log.Fatalln(err)
	}

	want := e2e.LoginLoginUser{
		FirstName: "James",
		LastName:  "Bond",
	}
	got := res.GetLogin()
	test.AssertEqual(want, got, t)
}
