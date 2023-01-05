package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetHourlyTotals(ctx context.Context, vars map[string]interface{}) (*HourlyTotals, error) {
	query := buildQueryHourlyTotals(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData HourlyTotals
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryHourlyTotals(vars map[string]interface{}) string {
	query := `query {
		hourlyTotals`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			timestamp
			totalTokens
			numberOfTrades
			orders
			settlements
			volumeUsd
			volumeEth
			feesUsd
			feesEth
			tokens {
			  id
			}
		}
	}`
	return query
}
