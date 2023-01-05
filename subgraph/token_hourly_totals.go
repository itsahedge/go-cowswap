package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetTokenHourlyTotals(ctx context.Context, vars map[string]interface{}) (*TokenHourlyTotals, error) {
	query := buildQueryTokenHourlyTotals(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData TokenHourlyTotals
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryTokenHourlyTotals(vars map[string]interface{}) string {
	query := `query {
		tokenHourlyTotals`
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
