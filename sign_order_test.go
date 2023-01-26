package go_cowswap

import (
	"strings"
	"testing"
)

func Test_SignatureSchemaEip712(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	sellToken := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	buyToken := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	sellAmount := "999275952760391053"
	buyAmount := "3238567244844690335887"
	order := &CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(address),
		SellAmount:        sellAmount,
		BuyAmount:         buyAmount,
		ValidTo:           uint32(1674752286),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         "886313204257255",
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     "eip712", // eip712 or ethsign
		From:              strings.ToLower(address),
	}

	// 3) Sign order using eip712 signature scheme
	order, err = client.SignOrder(order)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(order)
}

func Test_SignatureSchemeEthSign(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	sellToken := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	buyToken := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	sellAmount := "999275952760391053"
	buyAmount := "3238567244844690335887"
	order := &CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(address),
		SellAmount:        sellAmount,
		BuyAmount:         buyAmount,
		ValidTo:           uint32(1674752286),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         "886313204257255",
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     "ethsign", // eip712 or ethsign
		From:              strings.ToLower(address),
	}

	// 3) Sign order using eth_sign signature scheme
	order, err = client.SignOrder(order)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(order)
}

func Test_InvalidSignatureScheme(t *testing.T) {
	client, err := NewClient(Options)
	if err != nil {
		t.Fatal(err)
	}
	address := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	sellToken := "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	buyToken := "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	sellAmount := "999275952760391053"
	buyAmount := "3238567244844690335887"
	order := &CounterOrder{
		SellToken:         sellToken,
		BuyToken:          buyToken,
		Receiver:          strings.ToLower(address),
		SellAmount:        sellAmount,
		BuyAmount:         buyAmount,
		ValidTo:           uint32(1674752286),
		AppData:           "0x0000000000000000000000000000000000000000000000000000000000000000",
		Kind:              "sell",
		FeeAmount:         "886313204257255",
		PartiallyFillable: false,
		SellTokenBalance:  "erc20",
		BuyTokenBalance:   "erc20",
		SigningScheme:     "", // eip712 or ethsign
		From:              strings.ToLower(address),
	}

	// 3) Sign order using eth_sign signature scheme
	order, err = client.SignOrder(order)
	if err == nil {
		t.Error("want error for invalid signature")
	}
}
