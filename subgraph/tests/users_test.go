package tests

import (
	"context"
	"encoding/json"
	cowswap "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/subgraph"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	gql_client, err := subgraph.NewSubgraphClient(cowswap.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]interface{}{
		"first": 1,
		"block": map[string]interface{}{
			"number": 15114083,
		},
	}

	users, err := gql_client.GetUsers(context.Background(), vars)
	if err != nil {
		t.Fatal(err)
	}

	r, _ := json.MarshalIndent(users, "", "  ")
	t.Logf("%v", string(r))
}
