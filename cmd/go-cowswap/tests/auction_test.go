package go_cowswap_test

import (
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetAuction(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	res, err := client.GetAuction()
	if err != nil {
		t.Fatalf("GetAuction err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}

func TestClient_GetSolverAuctionById(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	res, err := client.GetSolverAuctionById(1)
	if err != nil {
		t.Fatalf("GetSolverAuctionById err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}

func TestClient_GetSolverAuctionByTxHash(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	txHash := ""
	res, err := client.GetSolverAuctionByTxHash(txHash)
	if err != nil {
		t.Fatalf("GetSolverAuctionByTxHash err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}
