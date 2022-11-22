package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

func (c *Client) Version(ctx context.Context) (*types.VersionResponse, int, error) {
	endpoint := "/version"
	var dataRes types.VersionResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
