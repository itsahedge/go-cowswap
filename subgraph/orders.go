package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetOrders(ctx context.Context, vars map[string]interface{}) (*Orders, error) {
	query := buildQueryOrders(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Orders
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryOrders(vars map[string]interface{}) string {
	query := `query {
		orders`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			tradesTimestamp
			invalidateTimestamp
			presignTimestamp
			isSigned
			isValid
			owner {
				id
				address
				firstTradeTimestamp
				isSolver
				numberOfTrades
				solvedAmountEth
				solvedAmountUsd
				tradedAmountEth
				tradedAmountUsd
			}
			trades {
				id
				timestamp
				gasPrice
				feeAmount
				txHash
				settlement {
					id
					txHash
					firstTradeTimestamp
					solver {
						id
					}
				}
				buyAmount
				sellAmount
				buyAmountUsd
				sellAmountUsd
				sellToken {
					id
			  	}
			}
		}
	}`
	return query
}
