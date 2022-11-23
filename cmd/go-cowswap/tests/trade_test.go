package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetTrades(t *testing.T) {
	client := go_cowswap.NewClient(util.Options)
	opts := &types.GetTrades{
		Owner: util.Options.EthAddress,
	}
	res, statusCode, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Fatalf("GetTrades err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
