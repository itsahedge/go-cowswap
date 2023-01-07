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
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
)

func main() {
	// Initialize the go-cowswap client on Goerli with default RPC
	options := cowswap.ConfigOpts{
		Network:    "goerli",
		Host:       cowswap.HostConfig["goerli"],
		RpcUrl:     cowswap.RpcConfig["goerli"],
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

---

### Get Quote

```go
package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
)

func main() {
	options := cowswap.ConfigOpts{
		Network:    "goerli",
		Host:       cowswap.HostConfig["goerli"],
		RpcUrl:     cowswap.RpcConfig["goerli"],
		EthAddress: "",
		PrivateKey: "",
	}
	client, err := cowswap.NewClient(options)
	o := &cowswap.QuoteReq{
		SellToken:           cowswap.TOKEN_ADDRESSES["goerli"]["WETH"],
		BuyToken:            cowswap.TOKEN_ADDRESSES["goerli"]["COW"],
		Receiver:            options.EthAddress,
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: "1000000000000000000",
		From:                options.EthAddress,
	}
	res, code, err := client.Quote(context.Background(), o)
	if err != nil {
		log.Fatalf("GetQuote err: %v", err)
	}
	fmt.Printf("statusCode: %v \n", code)
	fmt.Printf("%v \n", res)
}

```

### Fetch User Trades
```go
package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
)

func main() {
	client, err := cowswap.NewClient(cowswap.Options)
	owner := "0x.."
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
```


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
	options := cowswap.ConfigOpts{
		Network:    "goerli",
		Host:       cowswap.HostConfig["goerli"],
		RpcUrl:     cowswap.RpcConfig["goerli"],
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


### Placing an Order

```go
package main

import (
	"context"
	"fmt"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
	"strings"
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
	if client.TransactionSigner == nil {
		log.Fatalf("transaction signer was not initialized: %v", err)
	}
	ctx := context.Background()
	network := options.Network
	sellToken := cowswap.TOKEN_ADDRESSES[network]["WETH"]
	buyToken := cowswap.TOKEN_ADDRESSES[network]["COW"]
	sellAmountBeforeFee := "10000000000000000" // 0.01 ETH
	receiver := client.TransactionSigner.SignerPubKey.Hex()
	from := client.TransactionSigner.SignerPubKey.Hex()
	signingScheme := "eip712" // ethsign or eip712

	// 1) Fetch Order quote
	quoteReq := &cowswap.QuoteReq{
		SellToken:           sellToken,
		BuyToken:            buyToken,
		Receiver:            strings.ToLower(receiver),
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       signingScheme,
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: sellAmountBeforeFee,
		From:                strings.ToLower(from),
	}

	quoteResp, code, err := client.Quote(ctx, quoteReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status: %v \nresp: %v \n", code, quoteResp)

	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	validToFromQuote := quoteResp.Quote.ValidTo

	// 2) Build the Order
	order := &cowswap.CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(receiver),
		SellAmount:        sellAmountFromQuote,
		BuyAmount:         buyAmountFromQuote,
		ValidTo:           uint32(validToFromQuote),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         feeAmountFromQuote,
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     signingScheme,
		From:              strings.ToLower(from),
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		log.Fatalf("SignOrder err: %v", err)
	}

	// add the signature to the order:
	placed, code, err := client.CreateOrder(ctx, order)
	if err != nil {
		log.Fatalf("CreateOrderTest err: %v", err)
	}
	fmt.Printf("order placed: %v \n", placed)
}

```