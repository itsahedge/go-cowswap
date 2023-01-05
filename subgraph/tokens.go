package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetTokens(ctx context.Context, vars map[string]interface{}) (*Tokens, error) {
	query := buildQueryTokens(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Tokens
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryTokens(vars map[string]interface{}) string {
	query := `query {
		tokens`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			address
			firstTradeTimestamp
			name
			symbol
			decimals
			totalVolume
			priceEth
			priceUsd
			numberOfTrades
			totalVolumeUsd
			totalVolumeEth
		}
	}`
	return query
}
