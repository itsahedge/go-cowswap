package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

// TODO: handle responses for empty orders
func TestClient_GetOrdersByUid(t *testing.T) {
	client, err := go_cowswap2.NewClient(util.Options)
	uid := ""
	res, statusCode, err := client.GetOrdersByUid(context.Background(), uid)
	if err != nil {
		t.Fatalf("GetOrderByUid err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetOrdersByTxHash(t *testing.T) {
	client, err := go_cowswap2.NewClient(util.Options)
	txHash := ""
	res, statusCode, err := client.GetOrdersByTxHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("GetOrdersByTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetOrdersByUser(t *testing.T) {
	client, err := go_cowswap2.NewClient(util.Options)
	userAddress := util.Options.EthAddress
	opts := &go_cowswap2.OrdersPaginated{
		Limit:  "3",
		Offset: "1",
	}
	res, statusCode, err := client.GetOrdersByUser(context.Background(), userAddress, opts)
	if err != nil {
		t.Fatalf("GetOrdersByUser err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
