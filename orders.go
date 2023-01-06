package go_cowswap

import (
	"context"
	"fmt"
	"time"
)

type OrderByUidResponse struct {
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
}

// GetOrdersByUid - Get existing order from UID
func (c *Client) GetOrdersByUid(ctx context.Context, uid string) (*OrderByUidResponse, int, error) {
	if uid == "" {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_order_id", Description: "order UID not provided"}
	}
	endpoint := fmt.Sprintf("/orders/%s", uid)
	var dataRes OrderByUidResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil

}

type OrdersByTxHashResponse []struct {
	CreationDate                 time.Time   `json:"creationDate"`
	Owner                        string      `json:"owner"`
	UID                          string      `json:"uid"`
	AvailableBalance             interface{} `json:"availableBalance"`
	ExecutedBuyAmount            string      `json:"executedBuyAmount"`
	ExecutedSellAmount           string      `json:"executedSellAmount"`
	ExecutedSellAmountBeforeFees string      `json:"executedSellAmountBeforeFees"`
	ExecutedFeeAmount            string      `json:"executedFeeAmount"`
	Invalidated                  bool        `json:"invalidated"`
	Status                       string      `json:"status"`
	Class                        string      `json:"class"`
	SettlementContract           string      `json:"settlementContract"`
	FullFeeAmount                string      `json:"fullFeeAmount"`
	IsLiquidityOrder             bool        `json:"isLiquidityOrder"`
	SellToken                    string      `json:"sellToken"`
	BuyToken                     string      `json:"buyToken"`
	Receiver                     string      `json:"receiver"`
	SellAmount                   string      `json:"sellAmount"`
	BuyAmount                    string      `json:"buyAmount"`
	ValidTo                      int         `json:"validTo"`
	AppData                      string      `json:"appData"`
	FeeAmount                    string      `json:"feeAmount"`
	Kind                         string      `json:"kind"`
	PartiallyFillable            bool        `json:"partiallyFillable"`
	SellTokenBalance             string      `json:"sellTokenBalance"`
	BuyTokenBalance              string      `json:"buyTokenBalance"`
	SigningScheme                string      `json:"signingScheme"`
	Signature                    string      `json:"signature"`
	Interactions                 struct {
		Pre []interface{} `json:"pre"`
	} `json:"interactions"`
}

// GetOrdersByTxHash - Get orders by settlement transaction hash
func (c *Client) GetOrdersByTxHash(ctx context.Context, txHash string) (*OrdersByTxHashResponse, int, error) {
	if txHash == "" {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_tx_hash", Description: "transaction hash not provided"}
	}
	endpoint := fmt.Sprintf("%s/transactions/%s/orders", c.Host, txHash)
	var dataRes OrdersByTxHashResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}

type OrdersPaginated struct {
	Offset string
	Limit  string
}

type OrdersByUserResponse []struct {
	CreationDate                 time.Time   `json:"creationDate"`
	Owner                        string      `json:"owner"`
	UID                          string      `json:"uid"`
	AvailableBalance             interface{} `json:"availableBalance"`
	ExecutedBuyAmount            string      `json:"executedBuyAmount"`
	ExecutedSellAmount           string      `json:"executedSellAmount"`
	ExecutedSellAmountBeforeFees string      `json:"executedSellAmountBeforeFees"`
	ExecutedFeeAmount            string      `json:"executedFeeAmount"`
	Invalidated                  bool        `json:"invalidated"`
	Status                       string      `json:"status"`
	Class                        string      `json:"class"`
	SettlementContract           string      `json:"settlementContract"`
	FullFeeAmount                string      `json:"fullFeeAmount"`
	IsLiquidityOrder             bool        `json:"isLiquidityOrder"`
	SellToken                    string      `json:"sellToken"`
	BuyToken                     string      `json:"buyToken"`
	Receiver                     string      `json:"receiver"`
	SellAmount                   string      `json:"sellAmount"`
	BuyAmount                    string      `json:"buyAmount"`
	ValidTo                      int         `json:"validTo"`
	AppData                      string      `json:"appData"`
	FeeAmount                    string      `json:"feeAmount"`
	Kind                         string      `json:"kind"`
	PartiallyFillable            bool        `json:"partiallyFillable"`
	SellTokenBalance             string      `json:"sellTokenBalance"`
	BuyTokenBalance              string      `json:"buyTokenBalance"`
	SigningScheme                string      `json:"signingScheme"`
	Signature                    string      `json:"signature"`
	Interactions                 struct {
		Pre []interface{} `json:"pre"`
	} `json:"interactions"`
}

// GetOrdersByUser - Get orders of a user paginated
func (c *Client) GetOrdersByUser(ctx context.Context, userAddress string, opts *OrdersPaginated) (*OrdersByUserResponse, int, error) {
	if userAddress == "" {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_user_address", Description: "user address not provided"}
	}
	endpoint := fmt.Sprintf("/account/%s/orders", userAddress)
	var queries = make(map[string]interface{})
	if opts != nil {
		if opts.Limit != "" {
			queries["limit"] = opts.Limit
		}
		if opts.Offset != "" {
			queries["offset"] = opts.Offset
		}
	}
	var dataRes OrdersByUserResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}
