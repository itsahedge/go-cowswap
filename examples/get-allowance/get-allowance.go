package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
)

func main() {
	// Checking the allowance for an Address does not require private key to be set as its read-only
	network := "goerli"
	rpc := "https://eth-goerli-rpc.gateway.pokt.network"
	host := "https://api.cow.fi/goerli/api/v1"
	options := cowswap.ConfigOpts{
		Network: network,
		Host:    host,
		RpcUrl:  rpc,
	}
	client, err := cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Get allowance of User for WETH on CowSwap
	address := "0xcea7fb5b582c07129b8dc2fec4d4e5435b0968ff"
	tokenAddress := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	allowance, err := client.GetAllowance(ctx, address, tokenAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v token allowance: %v \n", tokenAddress, allowance)
}
