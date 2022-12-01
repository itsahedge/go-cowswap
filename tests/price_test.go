package go_cowswap

import (
	"context"
	"encoding/json"
	"github.com/itsahedge/go-cowswap"
	"github.com/itsahedge/go-cowswap/util"
	"testing"
)

func TestClient_GetNativePrice(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	token := util.TOKEN_ADDRESSES["mainnet"]["GNO"]
	res, statusCode, err := client.GetNativePrice(context.Background(), token)
	if err != nil {
		t.Fatalf("GetNativePrice err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
