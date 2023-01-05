package tests

import (
	"github.com/itsahedge/go-cowswap/subgraph"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestGqlClient_NewClient(t *testing.T) {
	client, err := subgraph.NewSubgraphClient(util.SUBGRAPH_MAINNET)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("client: %v", client)
}
