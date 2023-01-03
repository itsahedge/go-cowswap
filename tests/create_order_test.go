package go_cowswap

import (
	"context"
	"encoding/json"
	"fmt"
	go_cowswap2 "github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"strings"
	"testing"
)

func TestClient_CreateOrder(t *testing.T) {
	network := "goerli"
	client, err := go_cowswap2.NewClient(util.Options)
	if err != nil {
		t.Fatal(err)
	}

	sellToken := util.TOKEN_ADDRESSES[network]["WETH"]
	buyToken := util.TOKEN_ADDRESSES[network]["COW"]
	sellAmountBeforeFee := "10000000000000000" // 0.01 ETH
	receiver := client.TransactionSigner.SignerPubKey.Hex()
	from := client.TransactionSigner.SignerPubKey.Hex()

	// 1) Fetch Order quote
	quoteReq := &go_cowswap2.QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(receiver),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       util.SigningSchemeConfig[network], // ethsign or eip712
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

	//t.Log("values from quote i will be using in CreateOrder():")
	//t.Logf("SellAmount: %v ", sellAmountFromQuote)
	//t.Logf("BuyAmount: %v ", buyAmountFromQuote)
	//t.Logf("FeeAmount: %v ", feeAmountFromQuote)
	//t.Logf("AppData: %v ", appDataFromQuote)
	//t.Logf("ValidTo: %v ", validToFromQuote)

	// 2) Build the Order
	order := &go_cowswap2.CounterOrder{
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
		SigningScheme:     util.SigningSchemeConfig[network], // ethsign or eip712
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
	resp, code, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}
	uid := *resp
	t.Logf("order id: %v", uid)
}

func CreateOrderHandler(client *go_cowswap2.Client, network string) (string, error) {
	sellToken := util.TOKEN_ADDRESSES[network]["WETH"]
	buyToken := util.TOKEN_ADDRESSES[network]["COW"]
	seeAmountBeforeFee := "10000000000000000" // 0.01 ETH
	receiver := client.TransactionSigner.SignerPubKey.Hex()
	from := client.TransactionSigner.SignerPubKey.Hex()

	// 1) Fetch Order quote
	quoteReq := &go_cowswap2.QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(receiver),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       util.SigningSchemeConfig[network], // ethsign or eip712
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: seeAmountBeforeFee,
		From:                strings.ToLower(from),
	}

	quoteResp, code, err := client.Quote(context.Background(), quoteReq)
	if err != nil {
		return "", err
	}
	r, _ := json.MarshalIndent(quoteResp, "", "  ")
	fmt.Printf("statusCode: %v", code)
	fmt.Printf("%v", string(r))

	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	appDataFromQuote := quoteResp.Quote.AppData
	validToFromQuote := quoteResp.Quote.ValidTo

	fmt.Println("values from quote response to be used:")
	fmt.Printf("SellAmount: %v \n", sellAmountFromQuote)
	fmt.Printf("BuyAmount: %v \n", buyAmountFromQuote)
	fmt.Printf("FeeAmount: %v \n", feeAmountFromQuote)
	fmt.Printf("AppData: %v \n", appDataFromQuote)
	fmt.Printf("ValidTo: %v \n", validToFromQuote)

	// 2) Build the Order
	order := &go_cowswap2.CounterOrder{
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
		SigningScheme:     util.SigningSchemeConfig[network], // ethsign or eip712
		From:              strings.ToLower(from),
	}

	if client.TransactionSigner == nil {
		fmt.Printf("transaction signer was not initialized: %v", err)
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		fmt.Printf("SignOrder err: %v", err)
	}

	// 4) Place Trade order
	resp, code, err := client.CreateOrder(context.Background(), order)
	if err != nil {
		fmt.Print(err)
	}
	uid := *resp
	// {"errorType":"InsufficientFee","description":"Order does not include sufficient fee"}--- FAIL: TestCreateThenCancelOrder (1.36s)
	fmt.Println("======")
	fmt.Printf("order uid: %v \n", uid)
	fmt.Println("======")
	return uid, nil
}