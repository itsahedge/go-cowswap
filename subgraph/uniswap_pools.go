package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetUniswapPools(ctx context.Context, vars map[string]interface{}) (*UniswapPools, error) {
	query := buildQueryUniswapPools(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData UniswapPools
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryUniswapPools(vars map[string]interface{}) string {
	query := `query {
		uniswapPools`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			liquidity
			token0Price
			token1Price
			tick
			totalValueLockedToken0
			totalValueLockedToken1
			token0 {
			  id
			  address
			  name
			  symbol
			  decimals
			  priceEth
			  priceUsd
			}
			token1 {
			  id
			  address
			  name
			  symbol
			  decimals
			  priceEth
			  priceUsd
			}
		}
	}`
	return query
}
