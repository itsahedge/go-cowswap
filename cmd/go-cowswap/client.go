package go_cowswap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"math/big"
	"net/http"
)

type Client struct {
	Options util.ConfigOpts
	Network string
	Host    string

	Http             *http.Client
	Eip712OrderTypes apitypes.Types
	TypedDataDomain  apitypes.TypedDataDomain

	RpcUrl     string
	EthClient  *ethclient.Client
	ChainId    *big.Int
	ChainIdInt int

	TransactionSigner *TransactionSigner
}

func NewClient(options util.ConfigOpts) (*Client, error) {
	client := &Client{
		Options:          options,
		Http:             &http.Client{},
		Eip712OrderTypes: util.Eip712OrderTypes,
		TypedDataDomain:  util.Domain,
		Network:          options.Network,
		Host:             options.Host,
		RpcUrl:           options.RpcUrl,
	}
	if err := setClientNetwork(client, options); err != nil {
		return nil, err
	}
	if err := setClientRPC(client, options); err != nil {
		return nil, err
	}
	if err := setClientAuth(client, options); err != nil {
		return nil, err
	}
	return client, nil
}

func setClientNetwork(client *Client, options util.ConfigOpts) error {
	if options.Network != "" {
		client.Network = options.Network
		client.Host = util.HostConfig[options.Network]
		chainId := util.ChainIds[options.Network]
		client.ChainIdInt = chainId
		client.ChainId = big.NewInt(int64(chainId))
	}
	return nil
}

func setClientRPC(client *Client, options util.ConfigOpts) error {
	if options.RpcUrl != "" {
		client.RpcUrl = options.RpcUrl
		ethClient, err := ethclient.Dial(client.RpcUrl)
		if err != nil {
			return err
		}
		client.EthClient = ethClient
		chainId, err := ethClient.ChainID(context.Background())
		if err != nil {
			return err
		}
		client.ChainId = chainId
	}
	return nil
}

func setClientAuth(client *Client, options util.ConfigOpts) error {
	if options.PrivateKey != "" {
		transactionSigner, err := NewSigner(options.PrivateKey, client.ChainId)
		if err != nil {
			return fmt.Errorf("NewSigner err: %v\n", err)
		}
		client.TransactionSigner = transactionSigner
	}
	return nil
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("api err %d: %s", e.Code, e.Message)
}
