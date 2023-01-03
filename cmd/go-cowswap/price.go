package go_cowswap

import (
	"context"
	"fmt"
)

type NativePriceResponse struct {
	Price float64 `json:"price"`
}

func (c *Client) GetNativePrice(ctx context.Context, tokenAddress string) (*NativePriceResponse, int, error) {
	endpoint := fmt.Sprintf("/token/%s/native_price", tokenAddress)
	var dataRes NativePriceResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}
