package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestClient_GetVersion(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	res, code, err := client.Version(context.Background())
	if err != nil {
		t.Fatalf("Version err: %v", err)
	}
	t.Logf("statusCode: %v", code)
	t.Logf("%v", res)
}
