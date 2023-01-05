package tests

import (
	"github.com/itsahedge/go-cowswap/subgraph"
	"testing"
)

func TestGqlClient_NewClient(t *testing.T) {
	client, err := subgraph.NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("client: %v", client)
}
