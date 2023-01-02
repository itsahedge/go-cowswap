package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetOrdersByUid(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	uid := ""
	res, code, err := client.GetOrdersByUid(context.Background(), uid)
	if err != nil {
		t.Fatalf("GetOrderByUid err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}

func TestClient_GetOrdersByTxHash(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	txHash := ""
	res, code, err := client.GetOrdersByTxHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("GetOrdersByTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}

func TestClient_GetOrdersByUser(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	userAddress := util.Options.EthAddress
	opts := &go_cowswap.OrdersPaginated{
		Limit:  "3",
		Offset: "1",
	}
	res, code, err := client.GetOrdersByUser(context.Background(), userAddress, opts)
	if err != nil {
		t.Fatalf("GetOrdersByUser err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}
