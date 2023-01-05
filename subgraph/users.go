package subgraph

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

func (s *Client) GetUsers(ctx context.Context, vars map[string]interface{}) (*Users, error) {
	query := buildQuery(vars)
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	req.Header.Set("Cache-Control", "no-cache")
	var respData Users
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

func buildQuery(vars map[string]interface{}) string {
	query := `query {
		users`
	if len(vars) > 0 {
		query += ` (`
		query += buildVariables(vars) + ")"
	}
	query += ` {
			id
			address
			firstTradeTimestamp
			isSolver
			numberOfTrades
			solvedAmountEth
			solvedAmountUsd
			tradedAmountUsd
			tradedAmountEth
			ordersPlaced {
				id
				isSigned
			}
		}
	}`
	return query
}

// query filter: block, id, subgraph err
func (s *Client) GetUser(ctx context.Context, id string) (any, error) {
	req := graphql.NewRequest(`
		query ($id: String!) {
			user (id:$id) {
				id
				address
				firstTradeTimestamp
				isSolver
				numberOfTrades
				solvedAmountEth
				solvedAmountUsd
				tradedAmountUsd
				tradedAmountEth
				ordersPlaced {
					id
					isSigned
				}
			}
		}
	`)

	// wrap im handler
	req.Var("id", id)
	fmt.Println("req var::", req)

	//set header fields
	req.Header.Set("Cache-Control", "no-cache")

	var respData UserData
	if err := s.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}
