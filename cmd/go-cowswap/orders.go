package go_cowswap

import (
	"context"
	"errors"
	"fmt"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

func (c *Client) GetOrdersByUid(ctx context.Context, uid string) (*types.OrderByUidResponse, int, error) {
	if uid == "" {
		return nil, 404, errors.New("order UID not provided")
	}
	endpoint := fmt.Sprintf("/orders/%s", uid)
	var dataRes types.OrderByUidResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil || statusCode == 404 {
		return nil, statusCode, errors.New("order UID not found")
	}
	return &dataRes, statusCode, nil
}

func (c *Client) GetOrdersByTxHash(ctx context.Context, txHash string) (*types.OrdersByTxHashResponse, int, error) {
	if txHash == "" {
		return nil, 404, errors.New("transaction hash not provided")
	}
	endpoint := fmt.Sprintf("/transactions/%s/orders", txHash)
	var dataRes types.OrdersByTxHashResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil || statusCode == 404 {
		return nil, statusCode, errors.New("transaction hash not found")
	}
	return &dataRes, statusCode, nil
}

func (c *Client) GetOrdersByUser(ctx context.Context, userAddress string, opts *types.OrdersPaginated) (*types.OrdersByUserResponse, int, error) {
	if userAddress == "" {
		return nil, 404, errors.New("user address not provided")
	}
	endpoint := fmt.Sprintf("/account/%s/orders", userAddress)
	var queries = make(map[string]interface{})
	if opts != nil {
		if opts.Limit != "" {
			queries["limit"] = opts.Limit
		}
		if opts.Offset != "" {
			queries["offset"] = opts.Offset
		}
	}
	var dataRes types.OrdersByUserResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
