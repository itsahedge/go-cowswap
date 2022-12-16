package go_cowswap

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	"strings"
)

// SignOrder builds the CounterOrder from its Hash, Signs the Hash & Adds Signature
func (C *Client) SignOrder(order *CounterOrder) (*CounterOrder, error) {
	hash, err := C.Hash(order)
	if err != nil {
		return nil, fmt.Errorf("computing order hash: %v\n", err)
	}
	signatureBytes, err := C.SignHash(hash.Bytes(), C.TransactionSigner.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("signing order hash: %v\n", err)
	}
	order.Signature = fmt.Sprintf("0x%s", common.Bytes2Hex(signatureBytes))
	return order, nil
}

// Hash computes this counter order's hash.
func (C *Client) Hash(o *CounterOrder) (common.Hash, error) {
	var message = map[string]interface{}{
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

	domain := util.Domain
	domain.ChainId = math.NewHexOrDecimal256(C.ChainId.Int64())

	typedData := util.TypedData
	typedData.Domain = domain
	typedData.Message = message

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, fmt.Errorf("computing domain separator: %v\n", err)
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, fmt.Errorf("computing typed data hash: %v\n", err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return crypto.Keccak256Hash(rawData), nil
}

// SignHash sign the order hash with Transaction Signer Key
func (C *Client) SignHash(hash []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
	signatureBytes, err := crypto.Sign(accounts.TextHash(hash), pk)
	if err != nil {
		return nil, err
	}

	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}
	signatureBytes[64] = vParam
	return signatureBytes, nil
}
