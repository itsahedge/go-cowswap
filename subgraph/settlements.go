package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetSettlements(ctx context.Context, vars map[string]interface{}) (*Settlements, error) {
	query := buildQuerySettlements(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Settlements
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQuerySettlements(vars map[string]interface{}) string {
	query := `query {
		settlements`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			txHash
			firstTradeTimestamp
			solver {
				id
				address
			}
		}
	}`
	return query
}
