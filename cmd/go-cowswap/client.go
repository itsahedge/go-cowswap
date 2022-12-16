package go_cowswap

import (
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
	if options.Network != "" {
		client.Network = options.Network
		client.Host = util.HostConfig[options.Network]
		chainId := util.ChainIds[options.Network]
		client.ChainIdInt = chainId
		client.ChainId = big.NewInt(int64(chainId))
	}
	if options.RpcUrl != "" {
		client.RpcUrl = options.RpcUrl
		// add the eth client with rpc..
	}

	// change to withEthClient
	// if options.WithAuthEth != nil .. then add
	var err error
	client.EthClient, err = ethclient.Dial(client.RpcUrl)
	if err != nil {
		return nil, err
	}

	// WithAuth
	if options.PrivateKey != "" {
		client, err = client.WithAuth(options)
	}
	return client, nil
}

func (C *Client) WithAuth(options util.ConfigOpts) (*Client, error) {
	transactionSigner, err := NewSigner(options.PrivateKey, C.ChainId)
	if err != nil {
		return C, fmt.Errorf("NewSigner err: %v\n", err)
	}
	C.TransactionSigner = transactionSigner
	fmt.Println("WithAuth:", C.TransactionSigner)
	return C, nil
}

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("api err %d: %s", e.Code, e.Message)
}
