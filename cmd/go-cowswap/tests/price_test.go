package go_cowswap

import (
	"encoding/json"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetNativePrice(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	token := util.TOKEN_ADDRESSES["goerli"]["GNO"]
	res, err := client.GetNativePrice(token)
	if err != nil {
		t.Fatalf("GetNativePrice err: %v", err)
	}
	r, _ := json.MarshalIndent(res, "", "  ")
	t.Logf("%v", string(r))
}
