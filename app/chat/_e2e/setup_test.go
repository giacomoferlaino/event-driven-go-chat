package e2e

import (
	"chat/app/auth/config"
	"chat/app/auth/graph"
	"chat/pkg/env"
	"chat/pkg/test/e2esdk"
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
		Router:      graph.Router(),
		SeedData:    e2esdk.DefaultSeed(),
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
