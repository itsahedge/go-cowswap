package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"log"

	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
)

func main() {
	// Initialize the go-cowswap client on Goerli with default RPC
	options := util.ConfigOpts{
		Network:    "goerli",
		Host:       util.HostConfig["goerli"],
		RpcUrl:     util.RpcConfig["goerli"],
		EthAddress: "YOUR-ETHEREUM-ADDRESS",
		PrivateKey: "YOUR-PRIVATE-KEY",
	}
	client, err := cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// Fetch the Chain ID and Block Number from the Client
	chainId, err := client.EthClient.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	block, err := client.EthClient.BlockNumber(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("chaind ID: %v, block: %v", chainId, block)
}
