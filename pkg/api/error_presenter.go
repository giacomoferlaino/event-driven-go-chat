package api

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	graphqlDefaultErrorPresenter = graphql.DefaultErrorPresenter
)

func ErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphqlDefaultErrorPresenter(ctx, e)

	var httpErr *HttpError
	if errors.As(e, &httpErr) {
		err.Message = httpErr.Message
	} else {
		err.Message = NewInternalServerError(nil).Message
	}

	log.Println(e.Error())
	return err
}
