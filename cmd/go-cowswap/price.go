package go_cowswap

import (
	"context"
	"errors"
	"fmt"
)

type NativePriceResponse struct {
	Price float64 `json:"price"`
}

func (c *Client) GetNativePrice(ctx context.Context, tokenAddress string) (*NativePriceResponse, int, error) {
	if tokenAddress == "" {
		return nil, 404, errors.New("token address not provided")
	}
	endpoint := fmt.Sprintf("/token/%v/native_price", tokenAddress)
	var dataRes NativePriceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
