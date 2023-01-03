package go_cowswap

import (
	"context"
	"fmt"
)

type GetTrades struct {
	Owner    string
	OrderUid string
}
type TradesResponse []struct {
	BlockNumber          int    `json:"blockNumber"`
	LogIndex             int    `json:"logIndex"`
	OrderUID             string `json:"orderUid"`
	BuyAmount            string `json:"buyAmount"`
	SellAmount           string `json:"sellAmount"`
	SellAmountBeforeFees string `json:"sellAmountBeforeFees"`
	Owner                string `json:"owner"`
	BuyToken             string `json:"buyToken"`
	SellToken            string `json:"sellToken"`
	TxHash               string `json:"txHash"`
}

// GetTrades Exactly one of owner or order_uid has to be set.
func (c *Client) GetTrades(ctx context.Context, opts *GetTrades) (*TradesResponse, int, error) {
	endpoint := "/trades"
	if opts == nil {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_payload", Description: "must specify exactly one of owner or order_uid"}
	}
	if opts != nil {
		if opts.Owner != "" && opts.OrderUid != "" {
			return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_payload", Description: "must specify exactly one of owner or order_uid"}
		}
		if opts.Owner != "" {
			endpoint = fmt.Sprintf("%s?owner=%s", endpoint, opts.Owner)
		}
		if opts.OrderUid != "" {
			endpoint = fmt.Sprintf("%s?orderUid=%s", endpoint, opts.OrderUid)
		}
	}
	var dataRes TradesResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}
