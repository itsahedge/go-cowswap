package subgraph

import (
	"context"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"github.com/machinebox/graphql"
	"log"
	"testing"
)

func TestNew_GraphqlClient(t *testing.T) {
	gql_client := graphql.NewClient(util.SUBGRAPH_MAINNET)
	req := graphql.NewRequest(`
        {
			users(first: 5) {
				id
				address
				firstTradeTimestamp
				ordersPlaced {
				  id
				}
			}
		}
    `)

	// set any variables
	req.Var("key", "value")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData any
	if err := gql_client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	t.Log(respData)
}
