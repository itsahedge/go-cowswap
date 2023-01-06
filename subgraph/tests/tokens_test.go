package tests

import (
	"context"
	"encoding/json"
	cowswap "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/subgraph"
	"testing"
)

func Test_GetTokens(t *testing.T) {
	gql_client, err := subgraph.NewSubgraphClient(cowswap.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]interface{}{
		"orderBy":        "numberOfTrades",
		"orderDirection": "desc",
		"first":          2,
		//"block": map[string]interface{}{
		//	"number": 15114083,
		//},
	}

	resp, err := gql_client.GetTokens(context.Background(), vars)
	if err != nil {
		t.Fatal(err)
	}

	r, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%v", string(r))
}
