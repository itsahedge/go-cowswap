package go_cowswap

import (
	"context"
	"encoding/json"
	"testing"
)

func TestClient_GetAllowance(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	ownerAddress := Options.EthAddress
	tokenAddress := TOKEN_ADDRESSES["goerli"]["WETH"]
	allowance, err := client.GetAllowance(context.Background(), ownerAddress, tokenAddress)
	if err != nil {
		t.Fatalf("GetAllowance err: %v", err)
	}
	result, _ := json.MarshalIndent(allowance, "", "  ")
	t.Logf("%v token allowance: %v \n", tokenAddress, string(result))
}

func TestClient_SetAllowance(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	// First check the allowance for WETH
	tokenAddress := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	allowance, err := client.GetAllowance(ctx, Options.EthAddress, tokenAddress)
	if err != nil {
		t.Fatal(err)
	}
	// if token allowance: 0
	if len(allowance.Bits()) == 0 {
		t.Logf("%v token allowance is: %v. Please call Approve() \n", tokenAddress, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, tokenAddress, tokenAmount)
		if err != nil {
			t.Logf("%v", err)
		} else {
			t.Logf("tx hash: %v", setAllowanceTx.Hash())
		}
	} else {
		t.Logf("%v token allowance: %v \n", tokenAddress, allowance)
	}
}
