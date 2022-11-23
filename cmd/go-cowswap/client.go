package go_cowswap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Network string
	Host    string

	Http             *http.Client
	Eip712OrderTypes apitypes.Types
	TypedDataDomain  apitypes.TypedDataDomain

	RpcUrl       string
	EthKeySinger *util.EthKeySinger
}

func NewClient(options types.Options) *Client {
	client := &Client{
		Http:             &http.Client{},
		Eip712OrderTypes: eip712OrderTypes,
		TypedDataDomain:  domain,
		Network:          options.Network,
		Host:             options.Host,
	}

	if options.Network != "" {
		client.Network = options.Network
		client.Host = NetworkConfig[options.Network]
	}

	if options.RpcUrl != "" {
		client.RpcUrl = options.RpcUrl
	}

	if options.PrivateKey != "" {
		client.EthKeySinger = util.NewSigner(options.PrivateKey)
	}

	return client
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
	case 200:
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
