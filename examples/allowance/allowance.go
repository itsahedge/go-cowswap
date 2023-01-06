package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
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
	ctx := context.Background()

	// Get allowance of User for WETH on CowSwap
	tokenAddress := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	allowance, err := client.GetAllowance(ctx, options.EthAddress, tokenAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v token allowance: %v \n", tokenAddress, allowance)

	// Leave empty for Unlimited allowance
	tokenAmount := ""
	setAllowanceTx, err := client.SetAllowance(ctx, tokenAddress, tokenAmount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx hash: %v", setAllowanceTx.Hash())

}
