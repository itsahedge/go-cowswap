package test

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	cowswap "github.com/itsahedge/go-cowswap"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
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
		t.Fatal(err)
	}
	if client.TransactionSigner != nil {
		t.Logf("initialized client with a transaction signer: %v", client)
	} else {
		t.Logf("initialized client without a transaction signer: %v", client)
	}

	chainId, err := client.EthClient.ChainID(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", chainId)

	block, err := client.EthClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("current block: %v", block)
	addressList := cowswap.TOKEN_ADDRESSES[cowswap.Options.Network]
	for s, s2 := range addressList {
		t.Logf("%v, %v \n", s, s2)
	}
	resp, code, err := client.GetVersion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("version resp: %v", resp)
}
