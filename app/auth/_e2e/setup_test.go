package e2e

import (
	"chat/app/auth/_e2e/e2esdk"
	"chat/app/auth/config"
	"chat/app/auth/graph"
	"chat/pkg/env"
	"flag"
	"log"
	"testing"
)

var e2eEnv *e2esdk.Env

func TestMain(m *testing.M) {
	flag.Parse()
	err := env.Init("../.env")
	if err != nil {
		log.Panicln(err)
	}

	config := e2esdk.Config{
		Router: graph.Router(),
		KeycloakData: e2esdk.KeycloakData{
			Realm:  realm(),
			Client: client(),
			Users:  users(),
		},
		KeycloakUrl: config.KcUrl(),
	}

	e2eEnv, err = e2esdk.NewEnv(config)
	if err != nil {
		log.Panicln(err)
	}

	err = e2eEnv.Setup()
	defer e2eEnv.Teardown()
	if err != nil {
		log.Panicln(err)
	}

	m.Run()
}
