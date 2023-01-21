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
		panic(err)
	}
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
	res, code, err := client.Quote(context.Background(), quoteReq)
	if err != nil {
		log.Panicf("Quote err: %v", err)
	}
	fmt.Printf("statusCode: %v \n", code)
	fmt.Printf("%v \n", res)
}
