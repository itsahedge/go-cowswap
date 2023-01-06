package go_cowswap_test

import (
	"context"
	cowswap "github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client, err := cowswap.NewClient(cowswap.Options)
	opts := &cowswap.GetTrades{
		Owner: cowswap.Options.EthAddress,
	}
	res, code, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("res: %v", res)
}
