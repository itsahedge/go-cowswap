package go_cowswap_test

import (
	"context"
	"encoding/json"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestClient_GetQuote(t *testing.T) {
	client, err := go_cowswap2.NewClient(util.Options)
	o := &go_cowswap2.QuoteReq{
		SellToken:           util.TOKEN_ADDRESSES["mainnet"]["WETH"],
		BuyToken:            util.TOKEN_ADDRESSES["mainnet"]["COW"],
		Receiver:            util.Options.EthAddress,
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: "1000000000000000000",
		From:                util.Options.EthAddress,
	}
	res, statusCode, err := client.GetQuote(context.Background(), o)
	if err != nil {
		t.Fatalf("GetQuote err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
