package go_cowswap

import (
	"context"
	"errors"
	"fmt"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

func (c *Client) GetNativePrice(ctx context.Context, tokenAddress string) (*types.NativePriceResponse, int, error) {
	if tokenAddress == "" {
		return nil, 404, errors.New("token address not provided")
	}
	endpoint := fmt.Sprintf("/token/%v/native_price", tokenAddress)
	var dataRes types.NativePriceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
