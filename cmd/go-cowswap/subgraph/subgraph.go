package subgraph

import (
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"github.com/machinebox/graphql"
)

type SubgraphClient struct {
	GraphqlClient *graphql.Client
}

func NewSubgraphClient() (*SubgraphClient, error) {
	gql_client := graphql.NewClient(util.SUBGRAPH_MAINNET)
	client := &SubgraphClient{
		GraphqlClient: gql_client,
	}
	return client, nil
}
