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




The following example demonstrates how to use the `Version()` function to get the version of the Ethereum client:

```go
package main

import (
	"context"
	"fmt"
	"log"

	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
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
	client, err := go_cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}

	// Call the Version function to fetch /version
	version, code, err := client.Version(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status code: %v, version: %v", code, version)
}
```



