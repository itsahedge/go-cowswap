package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetTokenDailyTotals(ctx context.Context, vars map[string]interface{}) (*TokenDailyTotals, error) {
	query := buildQueryTokenDailyTotals(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData TokenDailyTotals
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryTokenDailyTotals(vars map[string]interface{}) string {
	query := `query {
		tokenDailyTotals`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			token {
			  id
			}
			timestamp
			totalVolume
			totalVolumeUsd
			totalVolumeEth
			totalTrades
			openPrice
			closePrice
			higherPrice
			lowerPrice
			averagePrice
		}
	}`
	return query
}
