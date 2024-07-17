package e2esdk

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Env struct {
	Router     *gin.Engine
	TestServer *httptest.Server
	GQLClient  *GQLClient
	Keycloak   *Keycloak
}

func NewEnv(config Config) (*Env, error) {
	testServer := httptest.NewServer(config.Router)
	gqlClient := NewGQLClient(testServer.URL)
	keycloack, err := NewKeycloak(config.KeycloakUrl, config.KeycloakData)
	if err != nil {
		return nil, err
	}

	return &Env{
		Router:     config.Router,
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
