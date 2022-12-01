package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client, err := go_cowswap2.NewClient(util.Options)
	opts := &go_cowswap2.GetTrades{
		Owner: util.Options.EthAddress,
	}
	res, statusCode, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
