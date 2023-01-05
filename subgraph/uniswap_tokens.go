package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetUniswapTokens(ctx context.Context, vars map[string]interface{}) (*UniswapTokens, error) {
	query := buildQueryUniswapTokens(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData UniswapTokens
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryUniswapTokens(vars map[string]interface{}) string {
	query := `query {
		uniswapTokens`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			address
			name
			symbol
			decimals
			priceEth
			priceUsd
		}
	}`
	return query
}
