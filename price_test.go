package go_cowswap

import (
	"context"
	"encoding/json"
	"testing"
)

func TestClient_GetNativePrice(t *testing.T) {
	client, err := NewClient(Options)
	token := TOKEN_ADDRESSES["goerli"]["GNO"]
	res, code, err := client.GetNativePrice(context.Background(), token)
	if err != nil {
		t.Fatalf("GetNativePrice err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("statusCode: %v", code)
	t.Logf("%v", string(r))
}
