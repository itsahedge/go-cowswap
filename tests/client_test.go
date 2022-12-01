package go_cowswap_test

import (
	"fmt"
	"github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestNewClient(t *testing.T) {
	network := "mainnet"
	options := util.ConfigOpts{
		Network: network,
		Host:    util.HostConfig[network],
		RpcUrl:  util.RpcConfig[network],
	}
	client, err := go_cowswap.NewClient(options)
	if err != nil {
		t.Fatal(err)
	}
	if client.TransactionSigner != nil {
		t.Logf("initialized client with a transaction signer: %v", client)
	} else {
		t.Logf("initialized client without a transaction signer: %v", client)
	}

	addressList := util.TOKEN_ADDRESSES[network]
	for s, s2 := range addressList {
		fmt.Printf("%v, %v \n", s, s2)
	}
}
