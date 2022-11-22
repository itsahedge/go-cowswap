package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

func (c *Client) GetQuote(ctx context.Context, o *types.QuoteReq) (*types.QuoteResponse, int, error) {
	endpoint := "/quote"
	var dataRes types.QuoteResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, o)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
