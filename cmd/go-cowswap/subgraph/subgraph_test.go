package subgraph

import (
	"context"
	"encoding/json"
	"testing"
)

func TestGqlClient_NewClient(t *testing.T) {
	client, err := NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("client: %v", client)
}

func TestGqlClient_Users(t *testing.T) {
	gql_client, err := NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]interface{}{
		"first": 1,
		//"block": map[string]interface{}{
		//	"number_gte": 15114083,
		//},
	}

	users, err := gql_client.GetUsers(context.Background(), vars)
	if err != nil {
		// handle error
	}

	r, _ := json.MarshalIndent(users, "", "  ")
	t.Logf("%v", string(r))
}
