package go_cowswap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/subgraph"
	"github.com/itsahedge/go-cowswap/util"
	"io/ioutil"
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

	Subgraph *subgraph.Client
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
		subgraph, err := subgraph.NewSubgraphClient(util.SubgraphConfig[options.Network])
		if err != nil {
			return err
		}
		client.Subgraph = subgraph
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

type ErrorCowResponse struct {
	Code        int    `json:"code"`
	ErrorType   string `json:"errorType"`
	Description string `json:"description"`
}

func (e *ErrorCowResponse) Error() string {
	return fmt.Sprintf("api err %d: %s, %s", e.Code, e.ErrorType, e.Description)
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
