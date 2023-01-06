package go_cowswap

import (
	"context"
)

// GetVersion - Information about the current deployed version of the API
func (c *Client) GetVersion(ctx context.Context) (*VersionResponse, int, error) {
	endpoint := "/version"
	var dataRes VersionResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, &ErrorCowResponse{Code: statusCode, ErrorType: "do_request_error", Description: err.Error()}
	}
	return &dataRes, statusCode, nil
}

type VersionResponse struct {
	Branch  string `json:"branch"`
	Commit  string `json:"commit"`
	Version string `json:"version"`
}
