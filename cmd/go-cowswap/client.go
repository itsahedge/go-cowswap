package go_cowswap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"io/ioutil"
	"math/big"
	"net/http"
)

type Client struct {
	Network string
	Host    string

	Http             *http.Client
	Eip712OrderTypes apitypes.Types
	TypedDataDomain  apitypes.TypedDataDomain

	RpcUrl    string
	EthClient *ethclient.Client
	ChainId   *big.Int

	TransactionSigner *TransactionSigner
}

func NewClient(options util.ConfigOpts) (*Client, error) {
	client := &Client{
		Http:             &http.Client{},
		Eip712OrderTypes: util.Eip712OrderTypes,
		TypedDataDomain:  util.Domain,
		Network:          options.Network,
		Host:             options.Host,
		RpcUrl:           options.RpcUrl,
	}
	if options.Network != "" {
		client.Network = options.Network
		client.Host = util.NetworkConfig[options.Network]
	}
	if options.RpcUrl != "" {
		client.RpcUrl = options.RpcUrl
	}

	var err error
	client.EthClient, err = ethclient.Dial(client.RpcUrl)
	if err != nil {
		return nil, err
	}
	chainId, err := client.EthClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	client.ChainId = chainId

	if options.PrivateKey != "" {
		transactionSigner, err := NewSigner(options.PrivateKey, chainId)
		if err != nil {
			return client, fmt.Errorf("NewSigner err: %v\n", err)
		}
		client.TransactionSigner = transactionSigner
	}
	return client, nil
}

func setQueryParam(endpoint *string, params []map[string]interface{}) {
	var first = true
	for _, param := range params {
		for i := range param {
			if first {
				*endpoint = fmt.Sprintf("%s?%s=%v", *endpoint, i, param[i])
				first = false
			} else {
				*endpoint = fmt.Sprintf("%s&%s=%v", *endpoint, i, param[i])
			}
		}
	}
}

func (c *Client) doRequest(ctx context.Context, endpoint, method string, expRes interface{}, reqData interface{}, opts ...map[string]interface{}) (int, error) {
	callURL := fmt.Sprintf("%s%s", c.Host, endpoint)

	var dataReq []byte
	var err error

	if reqData != nil {
		dataReq, err = json.Marshal(reqData)
		if err != nil {
			return 0, err
		}
	}

	if len(opts) > 0 && len(opts[0]) > 0 {
		setQueryParam(&callURL, opts)
	}

	req, err := http.NewRequestWithContext(ctx, method, callURL, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, err
	}

	resp, err := c.Http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200, 201:
		// note: CreateOrder returns 201..
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, err
			}
		}
		return resp.StatusCode, nil
	default:
		return resp.StatusCode, fmt.Errorf("%s", body)
	}
}
