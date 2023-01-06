package eth_sign

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func SignEthSign(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (sig []byte, err error) {
	hash, err := encodeSigning(typedData)
	if err != nil {
		return
	}
	// 2) SignHash - sign the order hash with Transaction Signer Key
	sig, err = signHash(hash.Bytes(), privateKey)
	if err != nil {
		return
	}
	return
}

func encodeSigning(typedData apitypes.TypedData) (hash common.Hash, err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash = crypto.Keccak256Hash(rawData)
	return
}

// SignHash - sign the order hash with Transaction Signer Key
func signHash(hash []byte, pk *ecdsa.PrivateKey) ([]byte, error) {
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
