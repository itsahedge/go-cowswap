package tests

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap/subgraph"
	"testing"
)

func Test_GetHourlyTotals(t *testing.T) {
	gql_client, err := subgraph.NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}

	//vars := map[string]interface{}{
	//	"orderBy":        "id",
	//	"orderDirection": "desc",
	//	"first":          5,
	//}

	resp, err := gql_client.GetHourlyTotals(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	r, _ := json.MarshalIndent(resp, "", "  ")
	t.Logf("%v", string(r))
}
