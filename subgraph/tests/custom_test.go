package tests

import (
	"context"
	"encoding/json"
	cowswap "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/subgraph"
	"testing"
)

func Test_GetCustomQuery(t *testing.T) {
	gql_client, err := subgraph.NewSubgraphClient(cowswap.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}
	customQuery := `
		query users {
		  users(where: {numberOfTrades_gt: 10, numberOfTrades_lt: 50 }) {
			id
			address
			firstTradeTimestamp
			isSolver
			numberOfTrades
			solvedAmountEth
			solvedAmountUsd
			tradedAmountUsd
			tradedAmountEth
			ordersPlaced {
				id
				isSigned
			}
		  }
		}
	`
	resp, err := gql_client.GetCustom(context.Background(), customQuery)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%v", string(r))
}
