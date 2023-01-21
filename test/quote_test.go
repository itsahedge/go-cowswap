package test

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/itsahedge/go-cowswap"
	"log"
	"testing"
)

func TestClient_GetQuote(t *testing.T) {
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
	options := go_cowswap.ConfigOpts{
		Network:    "goerli",
		Host:       go_cowswap.HostConfig["goerli"],
		RpcUrl:     go_cowswap.RpcConfig["goerli"],
		EthAddress: address,
		PrivateKey: privateKeyString,
	}
	client, err := go_cowswap.NewClient(options)
	o := &go_cowswap.QuoteReq{
		SellToken:           go_cowswap.TOKEN_ADDRESSES["goerli"]["WETH"],
		BuyToken:            go_cowswap.TOKEN_ADDRESSES["goerli"]["COW"],
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
		t.Fatalf("GetQuote err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("%v", res)
}
