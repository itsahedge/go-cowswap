package subgraph

import (
	"github.com/itsahedge/go-cowswap/util"
	"github.com/machinebox/graphql"
)

type Client struct {
	GraphqlClient *graphql.Client
}

func NewSubgraphClient() (*Client, error) {
	gql_client := graphql.NewClient(util.SUBGRAPH_MAINNET)
	client := &Client{
		GraphqlClient: gql_client,
	}
	return client, nil
}
