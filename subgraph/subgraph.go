package subgraph

import (
	"github.com/itsahedge/go-cowswap/util"
	"github.com/machinebox/graphql"
	gql "github.com/shurcooL/graphql"
)

type SubgraphClient struct {
	GraphqlClient *graphql.Client
	Gql           *gql.Client
}

func NewSubgraphClient() (*SubgraphClient, error) {
	gql_client := graphql.NewClient(util.SUBGRAPH_MAINNET)

	gql_client2 := gql.NewClient(util.SUBGRAPH_MAINNET, nil)

	client := &SubgraphClient{
		GraphqlClient: gql_client,
		Gql:           gql_client2,
	}
	return client, nil
}
