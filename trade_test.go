package go_cowswap

import (
	"context"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client, err := NewClient(Options)
	opts := &GetTrades{
		Owner: "0xcea7fb5b582c07129b8dc2fec4d4e5435b0968ff",
	}
	res, code, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("res: %v", res)
}
