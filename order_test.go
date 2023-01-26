package go_cowswap

import (
	"context"
	"encoding/json"
	"testing"
)

func TestClient_GetOrdersByUid(t *testing.T) {
	client, err := NewClient(Options)
	uid := "0x11bdfea6d0196279c372a58e1fa049e60d3812ce63ee37960d081e5f0e5004b875144248501e8629214cfdef09e1f0fe21bf83a563b8d8b2"
	res, code, err := client.GetOrdersByUid(context.Background(), uid)
	if err != nil {
		t.Fatalf("GetOrderByUid err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}

func TestClient_GetOrdersBySettlementTxHash(t *testing.T) {
	client, err := NewClient(Options)
	txHash := "0xa37a8f3d8bc60e75a40b967a53ef35f732eba8d6dba049d271593ccf58489d57"
	res, code, err := client.GetOrdersBySettlementTxHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("GetOrdersBySettlementTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}

func TestClient_GetOrdersByUser(t *testing.T) {
	client, err := NewClient(Options)
	userAddress := "0xbd4ad46efbddb7a0bd1d65df2e84698e5ce0bdd4"
	opts := &OrdersPaginated{
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
