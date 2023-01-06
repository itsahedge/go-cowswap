package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetPairs(ctx context.Context, vars map[string]interface{}) (*Pairs, error) {
	query := buildQueryPairs(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Pairs
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryPairs(vars map[string]interface{}) string {
	query := `query {
		pairs`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			token0 {
			  id
			}
			token1 {
			  id
			}
			token0Price
			token1Price
			token0relativePrice
			token1relativePrice
			volumeToken0
			volumeToken1
			volumeTradedEth
			volumeTradedUsd
		}
	}`
	return query
}
