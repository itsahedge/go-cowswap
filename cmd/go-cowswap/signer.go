package go_cowswap

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

type TransactionSigner struct {
	PrivateKey   *ecdsa.PrivateKey
	SignerPubKey common.Address
	Auth         *bind.TransactOpts
}

func NewSigner(privateKey string, chainId *big.Int) (*TransactionSigner, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	signerPubKey := crypto.PubkeyToAddress(*publicKeyECDSA)
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, err
	}
	signer := &TransactionSigner{
		PrivateKey:   pk,
		SignerPubKey: signerPubKey,
		Auth:         auth,
	}
	return signer, nil
}
