package graph

import (
	"chat/app/auth/graph/generated"
	"chat/pkg/api"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphHandler() gin.HandlerFunc {
	rootResolver := &Resolver{
		di: newDIContainer(),
	}
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rootResolver}))
	handler.SetErrorPresenter(api.ErrorPresenter)

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
