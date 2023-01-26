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


## Config

#### Before you can use the SDK, please set the ENV variables: `PRIVATE_KEY` & `ETH_ADDRESS` (refer to .env.example)


### Initializing the SDK requires config to be set
Config is of type `ConfigOpts`

#### Initializing with default config
To use default config, you only need to specify which network: `mainnet`, `goerli`, `xdai`

```go
// config.go
network := "goerli"
options := ConfigOpts{
    Network:    network,
    Host:       HostConfig[network],
    RpcUrl:     RpcConfig[network],
    EthAddress: os.Getenv("ETH_ADDRESS"),
    PrivateKey: os.Getenv("PRIVATE_KEY"),
}
```

#### Initializing with custom config
To use custom config, you only need to specify which network: `mainnet`, `goerli`, `xdai`

```go
// config.go
network := "goerli"
host := "https://api.cow.fi/goerli/api/v1"
customRpc := ""
options := ConfigOpts{
    Network:    network,
    Host:       host,
    RpcUrl:     customRpc,
    EthAddress: os.Getenv("ETH_ADDRESS"),
    PrivateKey: os.Getenv("PRIVATE_KEY"),
}
```

### The following example demonstrates how to initialize the SDK with Read & Write Functions:

```go
package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
)

func main() {
	// Initialize the go-cowswap client on Goerli 
	
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


### Testing

In order to run the tests properly:
- set your ENV variables for `PRIVATE_KEY` & `ETH_ADDRESS`.
- Set the network config to Goerli 
- have some Goerli ETH/WETH

```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```



