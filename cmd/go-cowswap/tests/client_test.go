package go_cowswap_test

import (
	"context"
	"fmt"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	if client.TransactionSigner != nil {
		fmt.Println("initialized client without a transaction signer")
	} else {
		fmt.Println("initialized client with a transaction signer")
	}
	res, statusCode, err := client.Version(context.Background())
	if err != nil {
		t.Fatalf("Version err: %v", err)
	}
	t.Logf("status code: %v, response: %v", statusCode, res)
}
