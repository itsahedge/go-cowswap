package subgraph

import (
	"github.com/itsahedge/go-cowswap"
	"github.com/machinebox/graphql"
)

type Client struct {
	GraphqlClient *graphql.Client
}

func NewSubgraphClient() (*Client, error) {
	gql_client := graphql.NewClient(go_cowswap.SUBGRAPH_MAINNET)
	client := &Client{
		GraphqlClient: gql_client,
	}
	return client, nil
}
