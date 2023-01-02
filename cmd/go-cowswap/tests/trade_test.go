package go_cowswap_test

import (
	"context"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	opts := &go_cowswap.GetTrades{
		Owner: util.Options.EthAddress,
	}
	res, code, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("res: %v", res)
}
