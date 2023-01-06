package go_cowswap_test

import (
	"context"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetQuote(t *testing.T) {
	client, err := go_cowswap2.NewClient(go_cowswap2.Options)
	o := &go_cowswap2.QuoteReq{
		SellToken:           go_cowswap2.TOKEN_ADDRESSES["goerli"]["WETH"],
		BuyToken:            go_cowswap2.TOKEN_ADDRESSES["goerli"]["COW"],
		Receiver:            go_cowswap2.Options.EthAddress,
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: "1000000000000000000",
		From:                go_cowswap2.Options.EthAddress,
	}
	res, code, err := client.Quote(context.Background(), o)
	if err != nil {
		t.Fatalf("GetQuote err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("%v", res)
}
