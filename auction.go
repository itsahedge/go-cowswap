package go_cowswap

import (
	"context"
	"fmt"
	"time"
)

type AuctionResponse struct {
	ID                    int `json:"id"`
	Block                 int `json:"block"`
	LatestSettlementBlock int `json:"latestSettlementBlock"`
	Orders                []struct {
		SellToken                    string    `json:"sellToken"`
		BuyToken                     string    `json:"buyToken"`
		Receiver                     string    `json:"receiver"`
		SellAmount                   string    `json:"sellAmount"`
		BuyAmount                    string    `json:"buyAmount"`
		ValidTo                      int       `json:"validTo"`
		AppData                      string    `json:"appData"`
		FeeAmount                    string    `json:"feeAmount"`
		Kind                         string    `json:"kind"`
		PartiallyFillable            bool      `json:"partiallyFillable"`
		SellTokenBalance             string    `json:"sellTokenBalance"`
		BuyTokenBalance              string    `json:"buyTokenBalance"`
		SigningScheme                string    `json:"signingScheme"`
		Signature                    string    `json:"signature"`
		From                         string    `json:"from"`
		QuoteID                      int       `json:"quoteId"`
		CreationTime                 time.Time `json:"creationTime"`
		Owner                        string    `json:"owner"`
		UID                          string    `json:"UID"`
		AvailableBalance             string    `json:"availableBalance"`
		ExecutedSellAmount           string    `json:"executedSellAmount"`
		ExecutedSellAmountBeforeFees string    `json:"executedSellAmountBeforeFees"`
		ExecutedBuyAmount            string    `json:"executedBuyAmount"`
		ExecutedFeeAmount            string    `json:"executedFeeAmount"`
		Invalidated                  bool      `json:"invalidated"`
		Status                       string    `json:"status"`
		FullFeeAmount                string    `json:"fullFeeAmount"`
		IsLiquidityOrder             bool      `json:"isLiquidityOrder"`
		EthflowData                  struct {
			IsRefunded  bool `json:"isRefunded"`
			UserValidTo int  `json:"userValidTo"`
		} `json:"ethflowData"`
		OnchainUser string `json:"onchainUser"`
	} `json:"orders"`
	Prices struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"prices"`
}

// GetAuction The current batch auction that solvers should be solving right now. Includes the list of solvable orders, the block on which the batch was created, as well as prices for all tokens being traded (used for objective value computation).
func (c *Client) GetAuction(ctx context.Context) (*AuctionResponse, int, error) {
	endpoint := "/auction"
	var dataRes AuctionResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}

type SolverAuctionByIdResponse struct {
	AuctionId                  int     `json:"auctionId"`
	TransactionHash            string  `json:"transactionHash"`
	GasPrice                   float64 `json:"gasPrice"`
	LiquidityCollectedBlock    int     `json:"liquidityCollectedBlock"`
	CompetitionSimulationBlock int     `json:"competitionSimulationBlock"`
	Solutions                  []struct {
		Solver    string `json:"solver"`
		Objective struct {
			Total   float64 `json:"total"`
			Surplus float64 `json:"surplus"`
			Fees    float64 `json:"fees"`
			Cost    float64 `json:"cost"`
			Gas     int     `json:"gas"`
		} `json:"objective"`
		Prices struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"prices"`
		Orders []struct {
			Id             string `json:"id"`
			ExecutedAmount string `json:"executedAmount"`
		} `json:"orders"`
		CallData string `json:"callData"`
	} `json:"solutions"`
}

// GetSolverAuctionById Returns the competition information by auction id.
func (c *Client) GetSolverAuctionById(ctx context.Context, auctionId int) (*SolverAuctionByIdResponse, int, error) {
	if auctionId < 0 {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_auction_id", Description: "invalid auction id"}
	}
	endpoint := fmt.Sprintf("/solver_competition/%v", auctionId)
	var dataRes SolverAuctionByIdResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}

type SolverAuctionByTxHashResponse struct {
	AuctionId                  int     `json:"auctionId"`
	TransactionHash            string  `json:"transactionHash"`
	GasPrice                   float64 `json:"gasPrice"`
	LiquidityCollectedBlock    int     `json:"liquidityCollectedBlock"`
	CompetitionSimulationBlock int     `json:"competitionSimulationBlock"`
	Solutions                  []struct {
		Solver    string `json:"solver"`
		Objective struct {
			Total   float64 `json:"total"`
			Surplus float64 `json:"surplus"`
			Fees    float64 `json:"fees"`
			Cost    float64 `json:"cost"`
			Gas     int     `json:"gas"`
		} `json:"objective"`
		Prices struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"prices"`
		Orders []struct {
			Id             string `json:"id"`
			ExecutedAmount string `json:"executedAmount"`
		} `json:"orders"`
		CallData string `json:"callData"`
	} `json:"solutions"`
}

// GetSolverAuctionByTxHash Returns the competition information by transaction hash.
func (c *Client) GetSolverAuctionByTxHash(ctx context.Context, txHash string) (*SolverAuctionByTxHashResponse, int, error) {
	if txHash == "" {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_tx_hash", Description: "invalid tx hash"}
	}
	endpoint := fmt.Sprintf("/solver_competition/by_tx_hash/%s", txHash)
	var dataRes SolverAuctionByTxHashResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}
