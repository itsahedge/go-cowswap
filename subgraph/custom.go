package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetCustom(ctx context.Context, q string) (any, error) {

	var respData any
	req := graphql.NewRequest(q)
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return respData, nil
}
