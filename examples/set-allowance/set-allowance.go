package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
)

func main() {
	// Setting the allowance for an Address requires private key to be set
	network := "goerli"
	rpc := "https://eth-goerli-rpc.gateway.pokt.network"
	host := "https://api.cow.fi/goerli/api/v1"
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	options := cowswap.ConfigOpts{
		Network:    network,
		Host:       host,
		RpcUrl:     rpc,
		EthAddress: address,
		PrivateKey: privateKeyString,
	}
	client, err := cowswap.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// First check the allowance for WETH
	tokenAddress := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	allowance, err := client.GetAllowance(ctx, address, tokenAddress)
	if err != nil {
		log.Fatal(err)
	}
	// if token allowance: 0
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
