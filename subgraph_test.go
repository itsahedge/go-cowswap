package go_cowswap

import (
	"context"
	"encoding/json"
	"testing"
)

func TestNewClient_Subgraph(t *testing.T) {
	client, err := NewClient(Options)
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
