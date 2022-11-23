package go_cowswap

import (
	"context"
	"errors"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

type GetTrades struct {
	Owner    string
	OrderUid string
}

// GetTrades Exactly one of owner or order_uid has to be set.
func (c *Client) GetTrades(ctx context.Context, opts *GetTrades) (*types.TradesResponse, int, error) {
	endpoint := "/trades"

	if opts == nil {
		return nil, 400, errors.New("Must specify exactly one of owner or order_uid.")
	}

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
