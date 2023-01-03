<p align="center">
<br />
<a href="#"><img src="https://cow.fi/images/logo-light.svg" width="200" alt=""/></a>
<br />
</p>
<h1 align="center">CowSwap Go SDK</h1>
<p align="center">
<a href="https://discord.com/invite/cowprotocol"><img alt="Join our Discord!" src="https://img.shields.io/discord/869166959739170836.svg?color=7289da&label=discord&logo=discord&style=flat"/></a>
</p>

# Installation

To install the SDK with the `go get` command, run the following:

```bash
go get github.com/itsahedge/go-cowswap
```

## Getting Started


### Instantiating the SDK


The following example demonstrates how to initialize the SDK with Read & Write Functions:

```go
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
```

### Checking & Setting Allowance

```go
package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"log"
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
```
