package go_cowswap

import (
	"context"
	"log"
	"strings"
	"testing"
)

func TestClient_CreateOrder(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	if client.TransactionSigner == nil || Options.PrivateKey == "" {
		t.Logf("transaction signer was not initialized properly with private key:\ntransaction signer: %v\nPrivateKey: %v", client.TransactionSigner, Options.PrivateKey)
		return
	}
	ctx := context.Background()

	// 1) Fetch Order quote to sell 0.01 ETH for COW
	sellToken := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	buyToken := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	sellAmountBeforeFee := "10000000000000000" // 0.01 ETH
	quoteReq := &QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(Options.EthAddress),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712", // eip712 or ethsign
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: sellAmountBeforeFee,
		From:                strings.ToLower(Options.EthAddress),
	}
	quoteResp, _, err := client.Quote(context.Background(), quoteReq)
	if err != nil {
		t.Logf("Quote err: %v", err)
	}

	// Check allowance for Sell Token
	allowance, err := client.GetAllowance(ctx, Options.EthAddress, sellToken)
	if err != nil {
		log.Fatal(err)
	}
	// if token allownace: 0
	if len(allowance.Bits()) == 0 {
		t.Logf("%v token allowance is: %v. Please call Approve() \n", sellToken, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, sellToken, tokenAmount)
		if err != nil {
			t.Logf("setting allowance err: %v", err)
			return
		} else {
			t.Logf("set token allowance tx hash: %v \n", setAllowanceTx.Hash())
			return
		}
	}

	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	//feeAmountFromQuote := quoteResp.Quote.FeeAmount
	validToFromQuote := quoteResp.Quote.ValidTo

	// 2) Build the Order
	order := &CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(Options.EthAddress),
		SellAmount:        sellAmountFromQuote,
		BuyAmount:         buyAmountFromQuote,
		ValidTo:           uint32(validToFromQuote),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         "9165489564161798",
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     "eip712", // eip712 or ethsign
		From:              strings.ToLower(Options.EthAddress),
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		t.Fatal(err)
	}
	createdOrder, code, err := client.CreateOrder(ctx, order)
	if err != nil {
		t.Fatalf("CreateOrderTest err: %v", err)
	}
	t.Logf("statusCode: %v \n", code)
	t.Logf("order created. order uid: %v \n", *createdOrder)
}
