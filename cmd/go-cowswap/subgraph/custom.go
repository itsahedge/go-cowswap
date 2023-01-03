package subgraph

import (
	"context"
	"github.com/machinebox/graphql"
)

func (c *SubgraphClient) GetCustom(ctx context.Context, q string) (any, error) {

	var respData any
	req := graphql.NewRequest(q)
	if err := c.GraphqlClient.Run(ctx, req, &respData); err != nil {
		return nil, err
	}
	return respData, nil
}
