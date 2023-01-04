package subgraph

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

type UserData struct {
	Users User `json:"user,omitempty"`
}

/* TODO
1) Get 1 User
2) Get Users
*/

// query filter: block, id, subgraph err
func (c *SubgraphClient) GetUser(ctx context.Context, id string) (any, error) {
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
	if err := c.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}

type UsersData struct {
	Users []User `json:"users,omitempty"`
}

type User struct {
	ID                  string         `json:"id,omitempty"`
	Address             string         `json:"address,omitempty"`
	FirstTradeTimestamp int64          `json:"firstTradeTimestamp,omitempty"`
	IsSolver            bool           `json:"isSolver,omitempty"`
	NumberOfTrades      int64          `json:"numberOfTrades,omitempty"`
	SolvedAmountEth     string         `json:"solvedAmountEth,omitempty"`
	SolvedAmountUsd     string         `json:"solvedAmountUsd,omitempty"`
	TradedAmountUsd     string         `json:"tradedAmountUsd,omitempty"`
	TradedAmountEth     string         `json:"tradedAmountEth,omitempty"`
	OrdersPlaced        []OrdersPlaced `json:"ordersPlaced,omitempty"`
}

type OrdersPlaced struct {
	ID       *string `json:"id,omitempty"`
	IsSigned *bool   `json:"isSigned,omitempty"`
}

func (c *SubgraphClient) GetUsers(ctx context.Context, vars map[string]interface{}) (*UsersData, error) {
	query := buildQuery(vars)
	req := graphql.NewRequest(query)

	// set any variables
	for key, value := range vars {
		req.Var(key, value)
	}

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	var respData UsersData
	if err := c.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}

	return &respData, nil
}

func buildQuery(vars map[string]interface{}) string {
	query := `query {
		users`
	if len(vars) > 0 {
		query += ` (`
	}
	var hasVars bool
	for key, value := range vars {
		if hasVars {
			query += `,`
		}
		if key == "block" {
			blockVars := value.(map[string]interface{})
			query += ` block: {`
			var hasBlockVars bool
			for blockKey, blockValue := range blockVars {
				if hasBlockVars {
					query += `,`
				}
				query += ` ` + blockKey + `:` + fmt.Sprint(blockValue)
				hasBlockVars = true
			}
			query += `}`
		} else {
			query += ` ` + key + `:` + fmt.Sprint(value)
		}
		hasVars = true
	}
	if len(vars) > 0 {
		query += `)`
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
