package go_cowswap

import (
	"context"
)

func (c *Client) CreateOrder(ctx context.Context, signedOrder *CounterOrder) (*string, int, error) {
	if c.TransactionSigner == nil {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "invalid_transaction_signer", Description: "invalid transaction signer"}
	}
	endpoint := "/orders"
	var dataRes string
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, signedOrder)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
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
