package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthKeySinger struct {
	PrivateKey *ecdsa.PrivateKey
}

func NewSigner(privateKey string) *EthKeySinger {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil
	}
	signer := &EthKeySinger{
		PrivateKey: pk,
	}
	return signer
}
