package go_cowswap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
	"testing"
)

var options = types.Options{
	Network:    "goerli",
	Host:       NetworkConfig["goerli"],
	RpcUrl:     "https://rpc.ankr.com/eth_goerli",
	EthAddress: "",
	PrivateKey: "",
}

func TestNewClient(t *testing.T) {
	client := NewClient(options)
	res, statusCode, err := client.Version(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("status code: %v, response: %v", statusCode, res)
}

func TestQuote(t *testing.T) {
	client := NewClient(options)
	o := &types.QuoteReq{
		SellToken:           GOERLI_WETH_TOKEN,
		BuyToken:            GOERLI_COW_TOKEN,
		Receiver:            "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: "1000000000000000000",
		From:                "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
	}
	res, statusCode, err := client.Quote(context.Background(), o)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
