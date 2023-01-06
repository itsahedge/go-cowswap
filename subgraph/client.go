package subgraph

import (
	"github.com/machinebox/graphql"
)

type Client struct {
	GraphqlClient *graphql.Client
}

func NewSubgraphClient(url string) (*Client, error) {
	gql_client := graphql.NewClient(url)
	client := &Client{
		GraphqlClient: gql_client,
	}
	return client, nil
}
