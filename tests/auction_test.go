package go_cowswap_test

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetAuction(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	res, code, err := client.GetAuction(context.Background())
	if err != nil {
		t.Fatalf("GetAuction err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}

func TestClient_GetSolverAuctionById(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	res, code, err := client.GetSolverAuctionById(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetSolverAuctionById err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}

func TestClient_GetSolverAuctionByTxHash(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	txHash := ""
	res, code, err := client.GetSolverAuctionByTxHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("GetSolverAuctionByTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}
