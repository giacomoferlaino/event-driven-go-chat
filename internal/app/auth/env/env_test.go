package env

import (
	"chat/internal/pkg/test"
	"errors"
	"flag"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		m.Run()
	}
}

func TestEnvInit(t *testing.T) {
	t.Run("if the env file is not found", func(t *testing.T) {
		t.Run("it should return an error", func(t *testing.T) {
			defer test.Stub(&godotenvLoad, func(filenames ...string) (err error) {
				return errors.New("Env file not found")
			})()

			err := Init()

			test.AssertNotEqual(err, nil, t)
		})
	})

	t.Run("it should return nil", func(t *testing.T) {
		defer test.Stub(&godotenvLoad, func(filenames ...string) (err error) {
			return nil
		})()

		defer test.Stub(&osGetenv, func(key string) string {
			return ""
		})()

		defer test.Stub(&ginSetMode, func(value string) {})()

		err := Init()

		test.AssertEqual(err, nil, t)
	})
}
