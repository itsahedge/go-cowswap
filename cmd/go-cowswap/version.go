package go_cowswap

import (
	"context"
	"errors"
)

type VersionResponse struct {
	Branch  string `json:"branch"`
	Commit  string `json:"commit"`
	Version string `json:"version"`
}

func (c *Client) Version(ctx context.Context) (*VersionResponse, int, error) {
	endpoint := "/version"
	var dataRes VersionResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, errors.New("not found")
	}
	return &dataRes, statusCode, nil
}
