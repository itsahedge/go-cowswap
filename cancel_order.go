package go_cowswap

import "C"
import (
	"context"
)

type CancelOrderReq struct {
	OrderUids     []string `json:"orderUids"`
	Signature     string   `json:"signature"`
	SigningScheme string   `json:"signingScheme"`
}

func (c *Client) CancelOrder(ctx context.Context, uid string) (*string, int, error) {
	endpoint := "/orders"
	signature, _, err := c.SignCancelOrder(uid)
	if err != nil {
		return nil, 404, &ErrorCowResponse{Code: 404, ErrorType: "sign_cancel_order_error", Description: err.Error()}
	}
	uids := []string{uid}
	reqPayload := &CancelOrderReq{
		Signature:     signature,
		OrderUids:     uids,
		SigningScheme: "eip712",
	}
	var dataRes string
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", &dataRes, reqPayload)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}
