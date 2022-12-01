package go_cowswap

import (
	"context"
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"strings"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	if err != nil {
		t.Fatal(err)
	}

	sellToken := util.WETH_TOKEN
	buyToken := util.USDC_TOKEN
	seeAmountBeforeFee := "100000000000000000" // 0.1 ETH
	receiver := client.TransactionSigner.SignerPubKey.Hex()
	from := client.TransactionSigner.SignerPubKey.Hex()

	// 1) Fetch Order quote
	quoteReq := &go_cowswap.QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(receiver),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "ethsign",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: seeAmountBeforeFee,
		From:                strings.ToLower(from),
	}

	quoteResp, statusCodeQuote, err := client.GetQuote(context.Background(), quoteReq)
	if err != nil {
		t.Fatal(err)
	}
	r, _ := json.MarshalIndent(quoteResp, "", "  ")
	t.Logf("status code: %v\n%v", statusCodeQuote, string(r))

	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	appDataFromQuote := quoteResp.Quote.AppData
	validToFromQuote := quoteResp.Quote.ValidTo

	t.Log("values from quote i will be using in CreateOrder():")
	t.Logf("SellAmount: %v ", sellAmountFromQuote)
	t.Logf("BuyAmount: %v ", buyAmountFromQuote)
	t.Logf("FeeAmount: %v ", feeAmountFromQuote)
	t.Logf("AppData: %v ", appDataFromQuote)
	t.Logf("ValidTo: %v ", validToFromQuote)

	// 2) Build the Order
	order := &go_cowswap.CounterOrder{
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
		SigningScheme:     "ethsign",
		From:              strings.ToLower(from),
	}

	if client.TransactionSigner == nil {
		t.Fatalf("transaction signer was not initialized: %v", err)
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		t.Fatalf("SignOrder err: %v", err)
	}

	// 4) Place Trade order
	resp, statusCode, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v\nresp: %v\n", statusCode, resp)
}
