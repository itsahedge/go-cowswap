package test

import (
	"context"
	"github.com/itsahedge/go-cowswap"
	"testing"
)

func TestClient_GetVersion(t *testing.T) {
	client, err := go_cowswap.NewClient(go_cowswap.Options)
	res, code, err := client.GetVersion(context.Background())
	if err != nil {
		t.Fatalf("Version err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("%v", res)
}
