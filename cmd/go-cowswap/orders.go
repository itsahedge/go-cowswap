package go_cowswap

import (
	"context"
	"errors"
	"fmt"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

func (c *Client) GetOrdersByUid(ctx context.Context, uid string) (*types.OrderByUidResponse, int, error) {
	endpoint := fmt.Sprintf("/orders/%s", uid)
	var dataRes types.OrderByUidResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil || statusCode == 404 {
		return nil, statusCode, errors.New("Order UID not found.")
	}
	return &dataRes, statusCode, nil
}

func (c *Client) GetOrdersByTxHash(ctx context.Context, txHash string) (*types.OrdersByTxHashResponse, int, error) {
	endpoint := fmt.Sprintf("/transactions/%s/orders", txHash)
	var dataRes types.OrdersByTxHashResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil || statusCode == 404 {
		return nil, statusCode, errors.New("Transaction Hash not found.")
	}
	return &dataRes, statusCode, nil
}
