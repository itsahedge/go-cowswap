package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetSolverAuctionById(t *testing.T) {
	client := go_cowswap.NewClient(util.Options)
	res, statusCode, err := client.GetSolverAuctionById(context.Background(), 100)
	if err != nil {
		t.Fatalf("GetSolverAuctionById err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
