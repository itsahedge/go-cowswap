package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
)

func main() {
	network := "goerli"
	rpc := "https://eth-goerli-rpc.gateway.pokt.network"
	host := "https://api.cow.fi/goerli/api/v1"
	options := cowswap.ConfigOpts{
		Network: network,
		Host:    host,
		RpcUrl:  rpc,
	}
	client, err := cowswap.NewClient(options)
	owner := "0xcea7fb5b582c07129b8dc2fec4d4e5435b0968ff"
	opts := &cowswap.GetTrades{
		Owner: owner,
	}
	res, code, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		fmt.Printf("GetTrades err: %v", err)
	}
	fmt.Printf("statusCode: %v \n", code)
	fmt.Printf("%v \n", res)
}
