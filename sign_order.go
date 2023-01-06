package go_cowswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/util/signature-scheme/eip712"
	"github.com/itsahedge/go-cowswap/util/signature-scheme/eth-sign"
	"strings"
)

type SignOrderFunc func(*Client, *CounterOrder) (string, *apitypes.TypedData, error)

var signOrderFuncs = map[string]SignOrderFunc{
	"ethsign": SignEthSignOrder,
	"eip712":  SignEip712Order,
}

func (c *Client) SignOrder(o *CounterOrder) (string, *apitypes.TypedData, error) {
	signFunc, ok := signOrderFuncs[o.SigningScheme]
	if !ok {
		return "", nil, fmt.Errorf("invalid signing scheme: %s", o.SigningScheme)
	}
	return signFunc(c, o)
}

// SignEthSignOrder Signs order with eth_sign signing scheme
func SignEthSignOrder(c *Client, o *CounterOrder) (string, *apitypes.TypedData, error) {
	message := map[string]interface{}{
		"sellToken":         o.SellToken,
		"buyToken":          o.BuyToken,
		"receiver":          o.Receiver,
		"sellAmount":        o.SellAmount,
		"buyAmount":         o.BuyAmount,
		"validTo":           fmt.Sprintf("%d", o.ValidTo),
		"appData":           common.Hex2Bytes(strings.TrimPrefix(o.AppData, "0x")),
		"feeAmount":         o.FeeAmount,
		"kind":              o.Kind,
		"partiallyFillable": o.PartiallyFillable,
		"sellTokenBalance":  o.SellTokenBalance,
		"buyTokenBalance":   o.BuyTokenBalance,
	}
	domain := apitypes.TypedDataDomain{
		Name:              "Gnosis Protocol",
		Version:           "v2",
		ChainId:           math.NewHexOrDecimal256(c.ChainId.Int64()), // NETWORK ID
		VerifyingContract: GPv2Settlement,
	}
	typedData := apitypes.TypedData{
		Types:       Eip712OrderTypes,
		PrimaryType: "Order",
		Domain:      domain,
		Message:     message,
	}

	sigBytes, err := eth_sign.SignEthSign(typedData, c.TransactionSigner.PrivateKey)
	if err != nil {
		return "", nil, err
	}
	signature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	return signature, &typedData, nil
}

// SignEip712Order Signs order with EIP712 signing scheme
func SignEip712Order(c *Client, o *CounterOrder) (string, *apitypes.TypedData, error) {
	message := map[string]interface{}{
		"sellToken":         o.SellToken,
		"buyToken":          o.BuyToken,
		"receiver":          o.Receiver,
		"sellAmount":        o.SellAmount,
		"buyAmount":         o.BuyAmount,
		"validTo":           fmt.Sprintf("%d", o.ValidTo),
		"appData":           common.Hex2Bytes(strings.TrimPrefix(o.AppData, "0x")),
		"feeAmount":         o.FeeAmount,
		"kind":              o.Kind,
		"partiallyFillable": o.PartiallyFillable,
		"sellTokenBalance":  o.SellTokenBalance,
		"buyTokenBalance":   o.BuyTokenBalance,
	}
	domain := apitypes.TypedDataDomain{
		Name:              "Gnosis Protocol",
		Version:           "v2",
		ChainId:           math.NewHexOrDecimal256(c.ChainId.Int64()),
		VerifyingContract: GPv2Settlement,
	}
	typedData := apitypes.TypedData{
		Types:       Eip712OrderTypes,
		PrimaryType: "Order",
		Domain:      domain,
		Message:     message,
	}

	sigBytes, err := eip712.SignTypedData(typedData, c.TransactionSigner.PrivateKey)
	if err != nil {
		return "", nil, err
	}
	orderSignature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	return orderSignature, &typedData, nil
}
