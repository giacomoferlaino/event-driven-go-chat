package test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

const (
	Path = "query"
)

type GQLClient struct {
	Ctx    context.Context
	Client graphql.Client
}

func NewGQLClient(baseUrl string) *GQLClient {
	url := fmt.Sprintf("%s/%s", baseUrl, Path)
	client := graphql.NewClient(url, http.DefaultClient)
	return &GQLClient{
		Ctx:    context.Background(),
		Client: client,
	}
}
