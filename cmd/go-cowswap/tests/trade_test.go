package go_cowswap_test

import (
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	opts := &go_cowswap.GetTrades{
		Owner: util.Options.EthAddress,
	}
	res, err := client.GetTrades(opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	t.Logf("res: %v", res)
}
