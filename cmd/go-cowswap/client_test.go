package go_cowswap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(util.Options)
	res, statusCode, err := client.Version(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("status code: %v, response: %v", statusCode, res)
}

func TestClient_GetQuote(t *testing.T) {
	client := NewClient(util.Options)
	o := &types.QuoteReq{
		SellToken:           util.WETH_TOKEN,
		BuyToken:            util.COW_TOKEN,
		Receiver:            "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		AppData:             "0x0000000000000000000000000000000000000000000000000000000000000000",
		PartiallyFillable:   false,
		SellTokenBalance:    "erc20",
		BuyTokenBalance:     "erc20",
		PriceQuality:        "fast",
		SigningScheme:       "eip712",
		OnchainOrder:        false,
		Kind:                "sell",
		SellAmountBeforeFee: "1000000000000000000",
		From:                "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
	}
	res, statusCode, err := client.GetQuote(context.Background(), o)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetAuction(t *testing.T) {
	client := NewClient(util.Options)
	res, statusCode, err := client.GetAuction(context.Background())
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetNativePrice(t *testing.T) {
	client := NewClient(util.Options)
	res, statusCode, err := client.GetNativePrice(context.Background(), util.GNO_TOKEN)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetSolverAuctionById(t *testing.T) {
	client := NewClient(util.Options)
	res, statusCode, err := client.GetSolverAuctionById(context.Background(), 100)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetTrades(t *testing.T) {
	client := NewClient(util.Options)
	opts := &GetTrades{
		Owner: "",
	}
	res, statusCode, err := client.GetTrades(context.Background(), opts)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetOrdersByUid(t *testing.T) {
	client := NewClient(util.Options)
	uid := ""
	res, statusCode, err := client.GetOrdersByUid(context.Background(), uid)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetOrdersByTxHash(t *testing.T) {
	client := NewClient(util.Options)
	txHash := ""
	res, statusCode, err := client.GetOrdersByTxHash(context.Background(), txHash)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}

func TestClient_GetOrdersByUser(t *testing.T) {
	client := NewClient(util.Options)
	userAddress := ""
	opts := &OrdersPaginated{
		Limit:  "3",
		Offset: "1",
	}
	res, statusCode, err := client.GetOrdersByUser(context.Background(), userAddress, opts)
	if err != nil {
		t.Error(err)
	}
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("status code: %v\n%v", statusCode, string(r))
}
