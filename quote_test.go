package go_cowswap

import (
	"context"
	"testing"
)

func QuoteTestBuilder(c *Client) (*QuoteResponse, error) {
	sellToken := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	buyToken := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	sellAmountBeforeFee := "10000000000000000" // 0.01 ETH
	o := &QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            Options.EthAddress,
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: sellAmountBeforeFee,
		From:                Options.EthAddress,
	}
	quoteResp, _, err := c.Quote(context.Background(), o)
	if err != nil {
		return nil, err
	}
	return quoteResp, nil
}

func TestClient_GetQuote(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	res, err := QuoteTestBuilder(client)
	t.Logf("%v", res)
}
