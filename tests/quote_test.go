package go_cowswap_test

import (
	"context"
	cowswap "github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetQuote(t *testing.T) {
	client, err := cowswap.NewClient(cowswap.Options)
	o := &cowswap.QuoteReq{
		SellToken:           cowswap.TOKEN_ADDRESSES["goerli"]["WETH"],
		BuyToken:            cowswap.TOKEN_ADDRESSES["goerli"]["COW"],
		Receiver:            cowswap.Options.EthAddress,
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: "1000000000000000000",
		From:                cowswap.Options.EthAddress,
	}
	res, code, err := client.Quote(context.Background(), o)
	if err != nil {
		t.Fatalf("GetQuote err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("%v", res)
}
