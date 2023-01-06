package go_cowswap_test

import (
	"context"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client, err := go_cowswap2.NewClient(go_cowswap2.Options)
	opts := &go_cowswap2.GetTrades{
		Owner: go_cowswap2.Options.EthAddress,
	}
	res, code, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("res: %v", res)
}
