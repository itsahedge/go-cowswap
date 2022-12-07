package go_cowswap

import (
	"context"
	"fmt"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func TestCreateThenCancelOrder(t *testing.T) {
	network := "goerli"
	options := util.ConfigOpts{
		Network:    network,
		Host:       util.HostConfig[network],
		RpcUrl:     util.RpcConfig[network],
		EthAddress: "",
		PrivateKey: "",
	}
	client, err := go_cowswap.NewClient(options)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client)

	// Create an Order that we'll cancel immediately
	uid, err := CreateOrderHandler(client, network)
	if err != nil {
		t.Fatal(err)
	}
	uidBytes := []byte(uid)
	fmt.Println("length of uidBytes:", len(uidBytes))

	// removes 0x
	//test, err := hex.DecodeString(uid[2:])
	//fmt.Println("length of test:", len(test))

	// Prepare the OrderCancellation payload
	order := &go_cowswap.CancelOrder{
		OrderUids:     uidBytes,
		SigningScheme: "eip712",
	}

	// TODO: fix Signing the cancelled order
	order, err = client.SignCancelOrder(order)
	if err != nil {
		t.Fatal("SignCancelOrder:", err)
	}
	fmt.Println("======")
	fmt.Println("order payload after SignCancelOrder():", order)
	fmt.Println("order.Signature::::", order.Signature)
	fmt.Println("======")

	// pass in the signed cancel order
	resp, statusCode, err := client.CancelOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v\nresp: %v\n", statusCode, resp)
}
