package go_cowswap

import (
	"context"
	"log"
	"strings"
	"testing"
)

func Test_CancelOrder(t *testing.T) {
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
	quoteResp, err := QuoteTestBuilder(client)
	if err != nil {
		t.Fatal(err)
	}
	sellToken := quoteResp.Quote.SellToken
	buyToken := quoteResp.Quote.BuyToken
	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	validToFromQuote := quoteResp.Quote.ValidTo

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
		FeeAmount:         feeAmountFromQuote,
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

	// Cancel the order
	uid := *createdOrder
	res, statusCode, err := client.CancelOrder(context.Background(), uid)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v", statusCode)
	t.Logf("res: %v", *res)
}

func Test_CancelOrders(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	if client.TransactionSigner == nil || Options.PrivateKey == "" {
		t.Logf("transaction signer was not initialized properly with private key:\ntransaction signer: %v\nPrivateKey: %v", client.TransactionSigner, Options.PrivateKey)
		return
	}
	ctx := context.Background()
	sellTokenIntent := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	// Check allowance for Sell Token
	allowance, err := client.GetAllowance(ctx, Options.EthAddress, sellTokenIntent)
	if err != nil {
		log.Fatal(err)
	}
	// if token allownace: 0
	if len(allowance.Bits()) == 0 {
		t.Logf("%v token allowance is: %v. Please call Approve() \n", sellTokenIntent, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, sellTokenIntent, tokenAmount)
		if err != nil {
			t.Logf("setting allowance err: %v", err)
			return
		} else {
			t.Logf("set token allowance tx hash: %v \n", setAllowanceTx.Hash())
			return
		}
	}

	// Place 2 Orders
	var uids []string
	for i := 0; i < 2; i++ {
		// Get quote
		quoteResp, err := QuoteTestBuilder(client)
		if err != nil {
			t.Fatal(err)
		}
		sellToken := quoteResp.Quote.SellToken
		buyToken := quoteResp.Quote.BuyToken
		sellAmountFromQuote := quoteResp.Quote.SellAmount
		buyAmountFromQuote := quoteResp.Quote.BuyAmount
		feeAmountFromQuote := quoteResp.Quote.FeeAmount
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
			FeeAmount:         feeAmountFromQuote,
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
		t.Logf("order #%v created. order uid: %v \n", i, *createdOrder)
		uids = append(uids, *createdOrder)
	}

	res, statusCode, err := client.CancelOrders(context.Background(), uids)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v", statusCode)
	t.Logf("res: %v", *res)
}
