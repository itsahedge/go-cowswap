package go_cowswap

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestCheckAllowance(t *testing.T) {
	client := go_cowswap.NewClient(util.Options)
	ownerAddress := util.Options.EthAddress
	tokenAddress := util.WETH_TOKEN
	allowance, err := client.GetAllowance(context.Background(), ownerAddress, tokenAddress)
	if err != nil {
		t.Fatalf("GetAllowance err: %v", err)
	}
	result, _ := json.MarshalIndent(allowance, "", "  ")
	t.Logf("%v token allowance: %v \n", tokenAddress, string(result))
}
