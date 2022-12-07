package go_cowswap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
)

func (c *Client) HashCancel(o *CancelOrder) (common.Hash, error) {
	var message = map[string]interface{}{
		"orderUids": o.OrderUids,
	}
	domain := util.Domain
	domain.ChainId = math.NewHexOrDecimal256(c.ChainId.Int64())
	typedData := util.TypedData
	typedData.Domain = domain
	typedData.Message = message

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, fmt.Errorf("computing domain separator: %v\n", err)
	}

	// Point to the OrderCancellations Data structure
	typedData.PrimaryType = "OrderCancellations"
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, fmt.Errorf("computing typed data hash: %v\n", err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return crypto.Keccak256Hash(rawData), nil
}

func (c *Client) SignCancelOrder(order *CancelOrder) (*CancelOrder, error) {
	hash, err := c.HashCancel(order)
	if err != nil {
		return nil, fmt.Errorf("computing order hash: %v\n", err)
	}
	signatureBytes, err := c.SignHash(hash.Bytes(), c.TransactionSigner.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("signing order hash: %v\n", err)
	}
	order.Signature = fmt.Sprintf("0x%s", common.Bytes2Hex(signatureBytes))
	return order, nil
}

type CancelOrder struct {
	OrderUids []byte `json:"order_uids"`
	Signature string `json:"signature"`
}

func (c *Client) CancelOrder(ctx context.Context, o *CancelOrder) (*string, int, error) {
	endpoint := "/orders"
	var dataRes *string
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", &dataRes, o)
	if err != nil {
		return nil, statusCode, err
	}
	return dataRes, statusCode, nil
}
