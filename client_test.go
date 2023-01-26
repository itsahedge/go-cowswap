package go_cowswap

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	options := Options
	client, err := NewClient(options)
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
	addressList := TOKEN_ADDRESSES[Options.Network]
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
