package graph

import (
	"chat/pkg/api"
	"chat/pkg/test"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestErrorPresenter(t *testing.T) {
	t.Run("if the error is of type HttpError", func(t *testing.T) {
		t.Run("it should return the specified error message", func(t *testing.T) {
			gqlError := &gqlerror.Error{}
			test.Stub(&graphqlDefaultErrorPresenter, func(ctx context.Context, err error) *gqlerror.Error {
				return gqlError
			})()

			httpError := api.NewHTTPError("test_error", http.StatusInternalServerError, nil)

			err := errorPresenter(context.Background(), httpError)

			test.AssertEqual(httpError.Message, err.Message, t)
		})
	})

	t.Run("if the error is not of type HttpError", func(t *testing.T) {
		t.Run("it should return a generic error message", func(t *testing.T) {
			gqlError := &gqlerror.Error{}
			test.Stub(&graphqlDefaultErrorPresenter, func(ctx context.Context, err error) *gqlerror.Error {
				return gqlError
			})()

			genericError := errors.New("not an HTTP error")

			err := errorPresenter(context.Background(), genericError)

			test.AssertEqual("internal_server_error", err.Message, t)
		})
	})
}
