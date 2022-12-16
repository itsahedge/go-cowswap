package go_cowswap

import (
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestClient_GetVersion(t *testing.T) {
	client, err := go_cowswap.NewClient(util.Options)
	res, err := client.Version()
	if err != nil {
		t.Fatalf("Version err: %v", err)
	}
	t.Logf("%v", res)
}
