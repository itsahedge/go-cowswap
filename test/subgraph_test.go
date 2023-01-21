package test

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap"
	"testing"
)

func TestNewClient_Subgraph(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Subgraph.GetTotals(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}
