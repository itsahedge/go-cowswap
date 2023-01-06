package go_cowswap

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetAllowance(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	ownerAddress := go_cowswap.Options.EthAddress
	tokenAddress := go_cowswap.TOKEN_ADDRESSES["goerli"]["COW"]
	allowance, err := client.GetAllowance(context.Background(), ownerAddress, tokenAddress)
	if err != nil {
		t.Fatalf("GetAllowance err: %v", err)
	}
	result, _ := json.MarshalIndent(allowance, "", "  ")
	t.Logf("%v token allowance: %v \n", tokenAddress, string(result))
}

func TestClient_SetAllowance(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount := "0"
	tokenToApprove := go_cowswap.TOKEN_ADDRESSES["goerli"]["COW"]
	tx, err := client.SetAllowance(context.Background(), tokenToApprove, tokenAmount)
	if err != nil {
		t.Fatalf("SetAllowance err : %v", err)
	}
	t.Logf("tx hash: %v", tx.Hash())
}
