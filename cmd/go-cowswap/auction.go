package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

// GetAuction The current batch auction that solvers should be solving right now. Includes the list of solvable orders, the block on which the batch was created, as well as prices for all tokens being traded (used for objective value computation).
func (c *Client) GetAuction(ctx context.Context) (*types.AuctionResponse, int, error) {
	endpoint := "/auction"
	var dataRes types.AuctionResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
