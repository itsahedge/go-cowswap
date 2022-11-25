package go_cowswap

import (
	"context"
	"errors"
)

type CreateOrderResponse struct {
	SellToken         string `json:"sellToken"`
	BuyToken          string `json:"buyToken"`
	Receiver          string `json:"receiver"`
	SellAmount        string `json:"sellAmount"`
	BuyAmount         string `json:"buyAmount"`
	ValidTo           int    `json:"validTo"`
	AppData           string `json:"appData"`
	FeeAmount         string `json:"feeAmount"`
	Kind              string `json:"kind"`
	PartiallyFillable bool   `json:"partiallyFillable"`
	SellTokenBalance  string `json:"sellTokenBalance"`
	BuyTokenBalance   string `json:"buyTokenBalance"`
	SigningScheme     string `json:"signingScheme"`
	Signature         string `json:"signature"`
	From              string `json:"from"`
	QuoteID           int    `json:"quoteId"`
}

// CounterOrder represents a Gnosis CounterOrder.
type CounterOrder struct {
	SellToken         string `json:"sellToken,omitempty"`
	BuyToken          string `json:"buyToken,omitempty"`
	Receiver          string `json:"receiver,omitempty"`
	SellAmount        string `json:"sellAmount"`
	BuyAmount         string `json:"buyAmount"`
	ValidTo           uint32 `json:"validTo,omitempty"`
	AppData           string `json:"appData,omitempty"`
	FeeAmount         string `json:"feeAmount"`
	Kind              string `json:"kind,omitempty"`
	PartiallyFillable bool   `json:"partiallyFillable"`
	Signature         string `json:"signature,omitempty"`
	SigningScheme     string `json:"signingScheme,omitempty"`
	SellTokenBalance  string `json:"sellTokenBalance,omitempty"`
	BuyTokenBalance   string `json:"buyTokenBalance,omitempty"`
	From              string `json:"from,omitempty"`
}

// Steps to Create an Order
// 1) Fetch Order quote
// 2) Build the Order
// 3) Sign the Order Hash
// 4) Send the Request => CreateOrder()

func (c *Client) CreateOrder(ctx context.Context, o *CounterOrder) (*CreateOrderResponse, int, error) {
	if c.TransactionSigner == nil {
		return nil, 404, errors.New("transaction signer was not initialized")
	}
	signedOrder, err := c.SignOrder(o)
	if err != nil {
		return nil, 404, err
	}
	endpoint := "/orders"
	var dataRes CreateOrderResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, signedOrder)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
