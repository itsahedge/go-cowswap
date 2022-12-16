package go_cowswap_test

import (
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetOrdersByUid(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	uid := ""
	res, err := client.GetOrdersByUid(uid)
	if err != nil {
		t.Fatalf("GetOrderByUid err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}

func TestClient_GetOrdersByTxHash(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	txHash := ""
	res, err := client.GetOrdersByTxHash(txHash)
	if err != nil {
		t.Fatalf("GetOrdersByTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}

func TestClient_GetOrdersByUser(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	userAddress := util.Options.EthAddress
	opts := &go_cowswap.OrdersPaginated{
		Limit:  "3",
		Offset: "1",
	}
	res, err := client.GetOrdersByUser(userAddress, opts)
	if err != nil {
		t.Fatalf("GetOrdersByUser err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}
