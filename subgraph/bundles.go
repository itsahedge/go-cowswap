package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetBundles(ctx context.Context, vars map[string]interface{}) (*Bundles, error) {
	query := buildQueryBundles(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Bundles
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryBundles(vars map[string]interface{}) string {
	query := `query {
		bundles`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			ethPriceUSD
		}
	}`
	return query
}
