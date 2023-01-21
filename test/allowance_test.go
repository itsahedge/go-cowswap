package test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	cowswap "github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetAllowance(t *testing.T) {
	client, err := cowswap.NewClient(cowswap.Options)
	if err != nil {
		t.Fatal(err)
	}
	ownerAddress := "0xcea7fb5b582c07129b8dc2fec4d4e5435b0968ff"
	tokenAddress := cowswap.TOKEN_ADDRESSES["goerli"]["WETH"]
	allowance, err := client.GetAllowance(context.Background(), ownerAddress, tokenAddress)
	if err != nil {
		t.Fatalf("GetAllowance err: %v", err)
	}
	result, _ := json.MarshalIndent(allowance, "", "  ")
	t.Logf("%v token allowance: %v \n", tokenAddress, string(result))
}

func TestClient_SetAllowance(t *testing.T) {
	network := "goerli"
	rpc := "https://eth-goerli-rpc.gateway.pokt.network"
	host := "https://api.cow.fi/goerli/api/v1"
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("error casting public key to ECDSA")
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
		t.Fatal(err)
	}
	ctx := context.Background()

	// First check the allowance for WETH
	tokenAddress := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	allowance, err := client.GetAllowance(ctx, address, tokenAddress)
	if err != nil {
		t.Fatal(err)
	}
	// if token allowance: 0
	if len(allowance.Bits()) == 0 {
		t.Logf("%v token allowance is: %v. Please call Approve() \n", tokenAddress, allowance)
		// if allowance is 0, set it.
		tokenAmount := ""
		setAllowanceTx, err := client.SetAllowance(ctx, tokenAddress, tokenAmount)
		if err != nil {
			t.Logf("%v", err)
		} else {
			t.Logf("tx hash: %v", setAllowanceTx.Hash())
		}
	} else {
		t.Logf("%v token allowance: %v \n", tokenAddress, allowance)
	}
}
