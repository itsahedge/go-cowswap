package test

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetNativePrice(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	token := go_cowswap.TOKEN_ADDRESSES["goerli"]["GNO"]
	res, code, err := client.GetNativePrice(context.Background(), token)
	if err != nil {
		t.Fatalf("GetNativePrice err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}
