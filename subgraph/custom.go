package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetCustom(ctx context.Context, query string) (any, error) {
	req := graphql.NewRequest(query)
	req.Header.Set("Cache-Control", "no-cache")
	var respData any
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}
