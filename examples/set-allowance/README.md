### Checking & Setting Allowance

```go
package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
)

func main() {
	// Initialize the go-cowswap client on Goerli with default RPC
	network := "goerli"
	rpc := "https://eth-goerli-rpc.gateway.pokt.network"
	host := "https://api.cow.fi/goerli/api/v1"
	privateKey := "YOUR-PRIVATE-KEY"
	address := "YOUR-ETHEREUM-ADDRESS"
	
	options := cowswap.ConfigOpts{
		Network:    network,
		Host:       host,
		RpcUrl:     rpc,
		EthAddress: address,
		PrivateKey: privateKey,
	}
	client, err := cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	// Get allowance of any User for WETH on CowSwap
	tokenAddress := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	allowance, err := client.GetAllowance(ctx, address, tokenAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v token allowance: %v \n", tokenAddress, allowance)

	if len(allowance.Bits()) == 0 {
		fmt.Printf("%v token allowance is: %v. Please call Approve() \n", tokenAddress, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, tokenAddress, tokenAmount)
		if err != nil {
			fmt.Printf("setting allowance err: %v", err)
		} else {
			fmt.Printf("tx hash: %v", setAllowanceTx.Hash())
		}
	} else {
		fmt.Printf("%v token allowance: %v \n", tokenAddress, allowance)
	}
}
```