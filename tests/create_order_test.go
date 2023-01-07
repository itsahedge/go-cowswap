package go_cowswap

import (
	"context"
	"encoding/json"
	cowswap "github.com/itsahedge/go-cowswap"
	"strings"
	"testing"
)

func TestClient_CreateOrder(t *testing.T) {
	network := "goerli"
	client, err := cowswap.NewClient(cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	if client.TransactionSigner == nil {
		t.Fatalf("transaction signer was not initialized: %v", err)
	}

	sellToken := cowswap.TOKEN_ADDRESSES[network]["WETH"]
	buyToken := cowswap.TOKEN_ADDRESSES[network]["COW"]
	sellAmountBeforeFee := "10000000000000000" // 0.01 ETH
	receiver := client.TransactionSigner.SignerPubKey.Hex()
	from := client.TransactionSigner.SignerPubKey.Hex()
	signingScheme := "eip712" // ethsign or eip712

	// 1) Fetch Order quote
	quoteReq := &cowswap.QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(receiver),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       signingScheme,
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: sellAmountBeforeFee,
		From:                strings.ToLower(from),
	}

	quoteResp, code, err := client.Quote(context.Background(), quoteReq)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := json.MarshalIndent(quoteResp, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))

	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	validToFromQuote := quoteResp.Quote.ValidTo

	// 2) Build the Order
	order := &cowswap.CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(receiver),
		SellAmount:        sellAmountFromQuote,
		BuyAmount:         buyAmountFromQuote,
		ValidTo:           uint32(validToFromQuote),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         feeAmountFromQuote,
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     signingScheme,
		From:              strings.ToLower(from),
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		t.Fatalf("SignOrder err: %v", err)
	}

	// 4) Place Trade order
	resp, code, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}
	uid := *resp
	t.Logf("order id: %v", uid)
}
