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
	if client.TransactionSigner == nil || options.PrivateKey == "" {
		log.Printf("transaction signer was not initialized properly with private key:\ntransaction signer: %v\nPrivateKey: %v", client.TransactionSigner, options.PrivateKey)
		return
	}
	ctx := context.Background()
    
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
	quoteResp, _, err := client.Quote(context.Background(), quoteReq)
	if err != nil {
		log.Panicf("Quote err: %v", err)
	}

	// Check allowance for Sell Token
	allowance, err := client.GetAllowance(ctx, address, sellToken)
	if err != nil {
		log.Fatal(err)
	}
	// if token allownace: 0
	if len(allowance.Bits()) == 0 {
		fmt.Printf("%v token allowance is: %v. Please call Approve() \n", sellToken, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, sellToken, tokenAmount)
		if err != nil {
			fmt.Printf("setting allowance err: %v", err)
			return
		} else {
			fmt.Printf("set token allowance tx hash: %v \n", setAllowanceTx.Hash())
			return
		}
	}
	
	sellAmountFromQuote := quoteResp.Quote.SellAmount
	buyAmountFromQuote := quoteResp.Quote.BuyAmount
	feeAmountFromQuote := quoteResp.Quote.FeeAmount
	validToFromQuote := quoteResp.Quote.ValidTo
	
	// 2) Build the Order
	order := &cowswap.CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(address),
		SellAmount:        sellAmountFromQuote,
		BuyAmount:         buyAmountFromQuote,
		ValidTo:           uint32(validToFromQuote),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         feeAmountFromQuote,
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     "eip712", // eip712 or ethsign
		From:              strings.ToLower(address),
	}

	// 3) Sign the order
	order, err = client.SignOrder(order)
	if err != nil {
		log.Fatal(err)
	}
	createdOrder, code, err := client.CreateOrder(ctx, order)
	if err != nil {
		log.Fatalf("CreateOrderTest err: %v", err)
	}
	fmt.Printf("statusCode: %v \n", code)
	fmt.Printf("order created. order uid: %v \n", *createdOrder)
}

```