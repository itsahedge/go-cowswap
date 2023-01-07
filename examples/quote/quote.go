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
