package go_cowswap

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestCheckAllowance(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	if err != nil {
		t.Fatal(err)
	}
	ownerAddress := util.Options.EthAddress
	tokenAddress := util.WETH_TOKEN
	allowance, err := client.GetAllowance(context.Background(), ownerAddress, tokenAddress)
	if err != nil {
		t.Fatalf("GetAllowance err: %v", err)
	}
	result, _ := json.MarshalIndent(allowance, "", "  ")
	t.Logf("%v token allowance: %v \n", tokenAddress, string(result))
}

func TestApproveSpender(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	if err != nil {
		t.Fatal(err)
	}
	tokenAmount := ""
	tokenToApprove := util.USDC_TOKEN
	tx, err := client.SetAllowance(context.Background(), tokenToApprove, tokenAmount)
	if err != nil {
		t.Fatalf("SetAllowance err : %v", err)
	}
	t.Logf("tx hash: %v", tx.Hash())
}
