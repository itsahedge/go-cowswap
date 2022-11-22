package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
	"testing"
)

var options = types.Options{
	Network:    "goerli",
	Host:       NetworkConfig["goerli"],
	RpcUrl:     "https://rpc.ankr.com/eth_goerli",
	EthAddress: "",
	PrivateKey: "",
}

func TestNewClient(t *testing.T) {
	client := NewClient(options)
	res, statusCode, err := client.Version(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("status code: %v, response: %v", statusCode, res)
}
