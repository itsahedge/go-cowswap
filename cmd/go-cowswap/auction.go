package go_cowswap

import (
	"context"
	"errors"
	"fmt"
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

// GetSolverAuctionById Returns the competition information by auction id.
func (c *Client) GetSolverAuctionById(ctx context.Context, auctionId int) (*types.SolverCompetitionResponse, int, error) {
	if auctionId == 0 {
		return nil, 404, errors.New("auction id not provided")
	}
	endpoint := fmt.Sprintf("/solver_competition/%v", auctionId)
	var dataRes types.SolverCompetitionResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

// GetSolverAuctionByTxHash Returns the competition information by transaction hash.
func (c *Client) GetSolverAuctionByTxHash(ctx context.Context, txHash string) (*types.SolverAuctionByTxHashResponse, int, error) {
	if txHash == "" {
		return nil, 404, errors.New("transaction hash not provided")
	}
	endpoint := fmt.Sprintf("/solver_competition/by_tx_hash/%v", txHash)
	var dataRes types.SolverAuctionByTxHashResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
