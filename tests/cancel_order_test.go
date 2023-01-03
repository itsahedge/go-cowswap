package go_cowswap

import (
	"context"
	"github.com/itsahedge/go-cowswap"
	util2 "github.com/itsahedge/go-cowswap/util"
	"testing"
)

func Test_VerifySignCancelOrder(t *testing.T) {
	client, err := go_cowswap.NewClient(util2.Options)
	if err != nil {
		t.Fatal(err)
	}
	uid := ""
	sig, typedData, err := client.SignCancelOrder(uid)
	if err != nil {
		t.Fatal(err)
	}
	////// CHECK SIGNATURE TO VERIFY OWNER
	hash, err := util2.EncodeForSigning(*typedData)
	if err != nil {
		t.Logf("encode for signing err: %v", err)
	}
	checkAddress := client.TransactionSigner.SignerPubKey
	isOwner := util2.VerifySig(checkAddress.Hex(), sig, hash.Bytes())
	t.Logf("order signature: %v", sig)
	t.Logf("typed data: %v", typedData)
	t.Logf("signature owner is verified: %v \n", isOwner)
}

func Test_CancelOrder(t *testing.T) {
	client, err := go_cowswap.NewClient(util2.Options)
	if err != nil {
		t.Fatal(err)
	}
	uid := ""
	res, statusCode, err := client.CancelOrder(context.Background(), uid)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("status code: %v", statusCode)
	t.Logf("res: %v", *res)
}
