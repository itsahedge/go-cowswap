package subgraph

import (
	"context"
	"encoding/json"
	"fmt"
	gql "github.com/shurcooL/graphql"
	"testing"
)

func TestGqlClient_NewClient(t *testing.T) {
	client, err := NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("client: %v", client)
}

func TestGqlClient_Users(t *testing.T) {
	gql_client, err := NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]interface{}{
		"first": 1,
		//"block": map[string]interface{}{
		//	"number_gte": 15114083,
		//},
	}

	users, err := gql_client.GetUsers(context.Background(), vars)
	if err != nil {
		// handle error
	}

	r, _ := json.MarshalIndent(users, "", "  ")
	t.Logf("%v", string(r))
}

type Query struct {
	Users []struct {
		ID                  gql.String
		Address             gql.String
		FirstTradeTimestamp gql.Int
		IsSolver            gql.Boolean
		NumberOfTrades      gql.Int
		SolvedAmountEth     gql.String
		SolvedAmountUsd     gql.String
		TradedAmountUsd     gql.String
		TradedAmountEth     gql.String
	} `graphql:"users(first: $first, block: { number_gte: $number_gte } )"`
}

type blockInput struct {
	NumberGte *gql.Int
}

func TestGql_Client2(t *testing.T) {
	client, err := NewSubgraphClient()
	if err != nil {
		t.Fatal(err)
	}

	variables := map[string]interface{}{
		"first":      gql.Int(10),
		"number_gte": gql.Int(16324534),
	}

	var q Query
	if err := client.Gql.Query(context.Background(), &q, variables); err != nil {
		fmt.Println(err)
	}
	for i, user := range q.Users {
		t.Logf("%v) %v", i, user.ID)
	}
}
