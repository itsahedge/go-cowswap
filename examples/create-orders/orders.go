package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
	"strings"
)

func main() {
	network := "goerli"
	rpc := "https://eth-goerli-rpc.gateway.pokt.network"
	host := "https://api.cow.fi/goerli/api/v1"
	// REPLACE PrivateKey & EthAddress with your own private key and address!
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	options := cowswap.ConfigOpts{
		Network:    network,
		Host:       host,
		RpcUrl:     rpc,
		EthAddress: address,
		PrivateKey: privateKeyString,
	}
	client, err := cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	if client.TransactionSigner == nil || options.PrivateKey == "" {
		log.Printf("transaction signer was not initialized properly with private key:\ntransaction signer: %v\nPrivateKey: %v", client.TransactionSigner, options.PrivateKey)
		return
	}
	ctx := context.Background()
	// 1) Fetch Order quote to sell 0.01 ETH for COW
	sellToken := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	buyToken := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	sellAmountBeforeFee := "10000000000000000" // 0.01 ETH
	quoteReq := &cowswap.QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(address),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712", // eip712 or ethsign
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: sellAmountBeforeFee,
		From:                strings.ToLower(address),
	}
	quoteResp, _, err := client.Quote(context.Background(), quoteReq)
	if err != nil {
		log.Panicf("Quote err: %v", err)
	}

	// Check allowance for Sell Token
	allowance, err := client.GetAllowance(ctx, address, sellToken)
	if err != nil {
		log.Fatal(err)
	}
	// if token allownace: 0
	if len(allowance.Bits()) == 0 {
		fmt.Printf("%v token allowance is: %v. Please call Approve() \n", sellToken, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, sellToken, tokenAmount)
		if err != nil {
			fmt.Printf("setting allowance err: %v", err)
			return
		} else {
			fmt.Printf("set token allowance tx hash: %v \n", setAllowanceTx.Hash())
			return
		}
	}
	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	validToFromQuote := quoteResp.Quote.ValidTo
	// 2) Build the Order
	order := &cowswap.CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(address),
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
		From:              strings.ToLower(address),
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		log.Fatal(err)
	}
	createdOrder, code, err := client.CreateOrder(ctx, order)
	if err != nil {
		log.Fatalf("CreateOrderTest err: %v", err)
	}
	fmt.Printf("statusCode: %v \n", code)
	fmt.Printf("order created. order uid: %v \n", *createdOrder)
}
