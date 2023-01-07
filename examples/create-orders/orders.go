package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
	"strings"
)

func main() {
	// Initialize the go-cowswap client on Goerli with default RPC
	options := cowswap.ConfigOpts{
		Network:    "goerli",
		Host:       cowswap.HostConfig["goerli"],
		RpcUrl:     cowswap.RpcConfig["goerli"],
		EthAddress: "",
		PrivateKey: "",
	}
	client, err := cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	if client.TransactionSigner == nil {
		log.Fatalf("transaction signer was not initialized: %v", err)
	}
	ctx := context.Background()
	network := options.Network
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

	quoteResp, code, err := client.Quote(ctx, quoteReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status: %v \nresp: %v \n", code, quoteResp)

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
		log.Fatal(err)
	}
	// add the signature to the order:
	//order.Signature = sig
	placed, code, err := client.CreateOrder(ctx, order)
	if err != nil {
		log.Fatalf("CreateOrderTest err: %v", err)
	}
	fmt.Printf("order placed: %v \n", placed)
}
