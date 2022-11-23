package go_cowswap

import (
	"context"
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetNativePrice(t *testing.T) {
	client := go_cowswap.NewClient(util.Options)
	res, statusCode, err := client.GetNativePrice(context.Background(), util.GNO_TOKEN)
	if err != nil {
		t.Fatalf("GetNativePrice err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
