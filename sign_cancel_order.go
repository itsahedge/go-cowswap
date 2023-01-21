package go_cowswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/util/signature-scheme/eip712"
)

// SignCancelOrder - Sign a single order for cancellation & generate the order signature
func (c *Client) SignCancelOrder(uid string) (string, *apitypes.TypedData, error) {
	var message = map[string]interface{}{
		"orderUids": []any{uid},
	}
	var domain = apitypes.TypedDataDomain{
		Name:              "Gnosis Protocol",
		Version:           "v2",
		ChainId:           math.NewHexOrDecimal256(c.ChainId.Int64()),
		VerifyingContract: GPv2Settlement,
	}
	TypedData.PrimaryType = "OrderCancellations"
	TypedData.Domain = domain
	TypedData.Message = message
	sigBytes, err := eip712.SignTypedData(TypedData, c.TransactionSigner.PrivateKey)
	if err != nil {
		return "", nil, err
	}
	orderSignature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	return orderSignature, &TypedData, nil
}

// SignCancelOrders - Sign multiple orders for cancellation & generate the order signature
func (c *Client) SignCancelOrders(uid []string) (string, *apitypes.TypedData, error) {
	var uids []any
	for _, id := range uid {
		uids = append(uids, id)
	}
	message := map[string]interface{}{
		"orderUids": uids,
	}
	var domain = apitypes.TypedDataDomain{
		Name:              "Gnosis Protocol",
		Version:           "v2",
		ChainId:           math.NewHexOrDecimal256(c.ChainId.Int64()),
		VerifyingContract: GPv2Settlement,
	}
	TypedData.PrimaryType = "OrderCancellations"
	TypedData.Domain = domain
	TypedData.Message = message
	sigBytes, err := eip712.SignTypedData(TypedData, c.TransactionSigner.PrivateKey)
	if err != nil {
		return "", nil, err
	}
	orderSignature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	return orderSignature, &TypedData, nil
}
