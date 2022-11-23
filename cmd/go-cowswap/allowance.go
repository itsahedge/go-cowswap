package go_cowswap

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/util"
	contract_binding "github.com/itsahedge/go-cowswap/pkg/contracts/generated"
	"math/big"
)

func (c *Client) GetAllowance(ctx context.Context, ownerAddress, tokenAddress string) (*big.Int, error) {
	if ownerAddress == "" {
		return nil, errors.New("Must provide an Owner Address.")
	}

	if tokenAddress == "" {
		return nil, errors.New("Must provide a Token Address.")
	}

	token := common.HexToAddress(tokenAddress)
	contract, err := contract_binding.NewErc20(token, c.EthClient)
	if err != nil {
		return nil, err
	}
	owner := common.HexToAddress(ownerAddress)
	spender := common.HexToAddress(util.GPv2_Vault_Relayer)
	allowance, err := contract.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
	if err != nil {
		return nil, err
	}
	return allowance, nil
}
