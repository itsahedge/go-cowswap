package go_cowswap

import (
	"context"
	"errors"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

// GetTrades Exactly one of owner or order_uid has to be set.
func (c *Client) GetTrades(ctx context.Context, opts *types.GetTrades) (*types.TradesResponse, int, error) {
	if opts == nil {
		return nil, 400, errors.New("must specify exactly one of owner or order_uid")
	}
	endpoint := "/trades"
	var queries = make(map[string]interface{})
	if opts != nil {
		if opts.Owner != "" && opts.OrderUid != "" {
			return nil, 400, errors.New("Must specify exactly one of owner or order_uid.")
		}
		if opts.Owner != "" {
			queries["owner"] = opts.Owner
		}
		if opts.OrderUid != "" {
			queries["orderUid"] = opts.OrderUid
		}
	}
	var dataRes types.TradesResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
