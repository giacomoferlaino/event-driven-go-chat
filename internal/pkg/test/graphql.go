package test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

const (
	Schema = "http"
	Host   = "127.0.0.1"
	Port   = "8080"
	Path   = "query"
)

type GQLClient struct {
	Ctx    context.Context
	Client graphql.Client
}

func NewGQLClient() GQLClient {
	url := fmt.Sprintf("%s://%s:%s/%s", Schema, Host, Port, Path)
	client := graphql.NewClient(url, http.DefaultClient)
	return GQLClient{
		Ctx:    context.Background(),
		Client: client,
	}
}
