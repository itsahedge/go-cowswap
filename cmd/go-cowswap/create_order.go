package go_cowswap

import (
	"context"
)

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

func (c *Client) CreateOrder(ctx context.Context, o *CounterOrder) (*string, int, error) {
	endpoint := "/orders"
	var dataRes string
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, o)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
