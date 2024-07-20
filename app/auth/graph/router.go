package graph

import (
	"chat/app/auth/graph/generated"
	"chat/pkg/api"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func errorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)

	var httpErr *api.HttpError
	if errors.As(e, &httpErr) {
		err.Message = e.Error()
	} else {
		err.Message = api.NewInternalServerError(nil).Message
	}

	return err
}

func graphHandler() gin.HandlerFunc {
	rootResolver := &Resolver{
		diContainer: newDIContainer(),
	}
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rootResolver}))
	handler.SetErrorPresenter(errorPresenter)

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	handler := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", playgroundHandler())
	router.POST("/query", graphHandler())

	return router
}
