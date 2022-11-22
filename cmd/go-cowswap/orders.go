package go_cowswap

import (
	"context"
	"errors"
	"fmt"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

func (c *Client) GetOrdersByUid(ctx context.Context, uid string) (*types.OrderByUidResponse, int, error) {
	endpoint := fmt.Sprintf("/orders/%s", uid)
	var dataRes types.OrderByUidResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil || statusCode == 400 {
		return nil, statusCode, errors.New("Order UID not found.")
	}
	return &dataRes, statusCode, nil
}
