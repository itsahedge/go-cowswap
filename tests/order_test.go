package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetOrdersByUid(t *testing.T) {
	client, err := go_cowswap2.NewClient(go_cowswap2.Options)
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
	client, err := go_cowswap2.NewClient(go_cowswap2.Options)
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
	client, err := go_cowswap2.NewClient(go_cowswap2.Options)
	userAddress := go_cowswap2.Options.EthAddress
	opts := &go_cowswap2.OrdersPaginated{
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
