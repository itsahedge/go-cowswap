package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetTotals(ctx context.Context, vars map[string]interface{}) (*Totals, error) {
	query := buildQueryTotals(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Totals
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryTotals(vars map[string]interface{}) string {
	query := `query {
		totals`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			tokens
			traders
			numberOfTrades
			settlements
			volumeUsd
			volumeEth
			feesUsd
			feesEth
		}
	}`
	return query
}
