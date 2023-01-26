package go_cowswap

import (
	"context"
	"testing"
)

func TestClient_GetVersion(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	res, code, err := client.GetVersion(context.Background())
	if err != nil {
		t.Fatalf("Version err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("%v", res)
}
