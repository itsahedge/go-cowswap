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