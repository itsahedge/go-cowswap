package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetTrades(ctx context.Context, vars map[string]interface{}) (*Trades, error) {
	query := buildQueryTrades(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Trades
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryTrades(vars map[string]interface{}) string {
	query := `query {
		trades`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			timestamp
			gasPrice
			feeAmount
			txHash
			settlement {
				id
				solver {
					id
				}
			}
			buyAmount
			sellAmount
			sellToken {
			  id
			}
			buyToken {
			  id
			}
			order {
				id
				owner {
					id
				}
			}
			buyAmountEth
			sellAmountEth
			buyAmountUsd
			sellAmountUsd
		}
	}`
	return query
}
