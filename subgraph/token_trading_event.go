package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (s *Client) GetTokenTradingEvents(ctx context.Context, vars map[string]interface{}) (*TokenTradingEvents, error) {
	query := buildQueryTokenTradingEvent(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData TokenTradingEvents
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQueryTokenTradingEvent(vars map[string]interface{}) string {
	query := `query {
		tokenTradingEvents`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			timestamp
			amountEth
			amountUsd
			token {
			  id
			}
			trade {
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
		}
	}`
	return query
}
