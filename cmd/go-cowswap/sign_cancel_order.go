package go_cowswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
)

// SignCancelOrder generates the order signature
func (c *Client) SignCancelOrder(uid string) (string, *apitypes.TypedData, error) {
	var message = map[string]interface{}{
		"orderUids": []any{uid},
	}
	var domain = apitypes.TypedDataDomain{
		Name:              "Gnosis Protocol",
		Version:           "v2",
		ChainId:           math.NewHexOrDecimal256(c.ChainId.Int64()),
		VerifyingContract: util.GPv2Settlement,
	}
	util.TypedData.PrimaryType = "OrderCancellations"
	util.TypedData.Domain = domain
	util.TypedData.Message = message
	sigBytes, err := util.SignTypedData(util.TypedData, c.TransactionSigner.PrivateKey)
	if err != nil {
		return "", nil, err
	}
	orderSignature := fmt.Sprintf("0x%s", common.Bytes2Hex(sigBytes))
	return orderSignature, &util.TypedData, nil
}
