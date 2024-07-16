package env

import (
	"chat/pkg/test"
	"errors"
	"testing"
)

func TestInit(t *testing.T) {
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

func TestPort(t *testing.T) {
	t.Run("if the PORT environment variable is not set", func(t *testing.T) {
		t.Run("it should return default port value", func(t *testing.T) {
			defer test.Stub(&osGetenv, func(key string) string {
				return ""
			})()

			want := defaultPort
			got := Port()

			test.AssertEqual(want, got, t)
		})
	})

	t.Run("if the PORT environment variable is set", func(t *testing.T) {
		t.Run("it should return the environment variable value", func(t *testing.T) {
			want := "9090"
			defer test.Stub(&osGetenv, func(key string) string {
				return want
			})()

			got := Port()

			test.AssertEqual(want, got, t)
		})
	})
}
