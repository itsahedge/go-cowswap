package go_cowswap_test

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestClient_GetAuction(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	res, statusCode, err := client.GetAuction(context.Background())
	if err != nil {
		t.Fatalf("GetAuction err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetSolverAuctionById(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	res, statusCode, err := client.GetSolverAuctionById(context.Background(), 130)
	if err != nil {
		t.Fatalf("GetSolverAuctionById err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetSolverAuctionByTxHash(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	txHash := ""
	res, statusCode, err := client.GetSolverAuctionByTxHash(context.Background(), txHash)
	if err != nil {
		t.Fatalf("GetSolverAuctionByTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
