package go_cowswap

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	go_cowswap "github.com/itsahedge/go-cowswap/cmd/go-cowswap"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"testing"
)

func NewSignerTest() (*ecdsa.PrivateKey, error) {
	privateKey := ""
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func Test_CancelOrder4(t *testing.T) {
	// manually get a random uid
	uid := "0xe9ff528b465b9b6419eee7d79ec5730856d96b09d0ec5473b972f521e3bb1a7475144248501e8629214cfdef09e1f0fe21bf83a563a3729a"
	checkAddress := common.HexToAddress("")
	order := &go_cowswap.CancelOrder{
		OrderUidsStr: uid,
	}
	var message = map[string]interface{}{
		"orderUids": []any{order.OrderUidsStr},
	}
	var Eip712OrderTypesTest = apitypes.Types{
		"EIP712Domain": {
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "version",
				Type: "string",
			},
			{
				Name: "chainId",
				Type: "uint256",
			},
			{
				Name: "verifyingContract",
				Type: "address",
			},
		},
		"Order": {
			{
				Name: "sellToken",
				Type: "address",
			},
			{
				Name: "buyToken",
				Type: "address",
			},
			{
				Name: "receiver",
				Type: "address",
			},
			{
				Name: "sellAmount",
				Type: "uint256",
			},
			{
				Name: "buyAmount",
				Type: "uint256",
			},
			{
				Name: "validTo",
				Type: "uint32",
			},
			{
				Name: "appData",
				Type: "bytes32",
			},
			{
				Name: "feeAmount",
				Type: "uint256",
			},
			{
				Name: "kind",
				Type: "string",
			},
			{
				Name: "partiallyFillable",
				Type: "bool",
			},
			{
				Name: "sellTokenBalance",
				Type: "string",
			},
			{
				Name: "buyTokenBalance",
				Type: "string",
			},
		},
		"OrderCancellations": {
			{
				Name: "orderUids",
				Type: "bytes[]",
			},
		},
	}
	var domain = apitypes.TypedDataDomain{
		Name:              "Gnosis Protocol",
		Version:           "v2",
		ChainId:           math.NewHexOrDecimal256(5),
		VerifyingContract: "0x9008D19f58AAbD9eD0D60971565AA8510560ab41",
	}
	var typedData = apitypes.TypedData{
		Types:       Eip712OrderTypesTest,
		PrimaryType: "OrderCancellations",
		Domain:      domain,
		Message:     message,
	}
	privateKey, err := NewSignerTest()
	if err != nil {
		t.Fatal(err)
	}

	sigBytes, err := util.SignTypedData(typedData, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	hash, err := util.EncodeForSigning(typedData)

	t.Logf("signed data (bytes): %v", sigBytes)
	orderSignature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	t.Logf("order signature: %v", orderSignature)

	t.Logf("%v", util.VerifySigTest(
		checkAddress.Hex(),
		orderSignature,
		hash.Bytes(),
	))
}
