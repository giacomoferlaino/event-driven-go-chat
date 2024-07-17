package test

import (
	"chat/app/auth/httpserver"
	"chat/pkg/keycloak/keycloacktest"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Env struct {
	Router     *gin.Engine
	TestServer *httptest.Server
	GQLClient  *GQLClient
	Keycloak   *keycloacktest.KeycloakTest
}

func NewEnv() (*Env, error) {
	router := httpserver.Router()
	testServer := httptest.NewServer(router)
	gqlClient := NewGQLClient(testServer.URL)
	keycloack, err := keycloacktest.New()
	if err != nil {
		return nil, err
	}

	return &Env{
		Router:     router,
		TestServer: testServer,
		GQLClient:  gqlClient,
		Keycloak:   keycloack,
	}, nil
}

func (e *Env) Setup() error {
	err := e.Keycloak.Setup()
	if err != nil {
		return err
	}
	return nil
}

func (e *Env) Teardown() {
	e.Keycloak.Teardown()
	e.TestServer.Close()
}
