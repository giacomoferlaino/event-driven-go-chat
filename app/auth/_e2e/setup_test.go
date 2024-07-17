package e2e

import (
	"chat/pkg/env"
	"chat/pkg/test"
	"flag"
	"log"
	"testing"
)

var testEnv *test.Env

func TestMain(m *testing.M) {
	flag.Parse()
	env.Init()
	var err error
	testEnv, err = test.NewEnv()
	if err != nil {
		log.Panicln(err)
	}

	err = testEnv.Setup()
	defer testEnv.Teardown()
	if err != nil {
		log.Panicln(err)
	}

	m.Run()
}
