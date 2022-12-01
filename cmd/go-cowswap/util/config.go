package util

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type ConfigOpts struct {
	Network string
	Host    string
	RpcUrl  string

	EthAddress string
	PrivateKey string
}

var Options = ConfigOpts{
	Network:    "mainnet",
	Host:       HostConfig["mainnet"],
	RpcUrl:     RpcConfig["mainnet"],
	EthAddress: "",
	PrivateKey: "",
}

var TOKEN_ADDRESSES = map[string]map[string]string{
	"mainnet": ETHEREUM_TOKEN_LIST,
	"goerli":  GOERLI_TOKEN_LIST,
	"xdai":    GNOSIS_SCAN_TOKEN_LIST,
}

var HostConfig = map[string]string{
	"mainnet": MAINNET_API,
	"goerli":  GOERLI_API,
	"xdai":    GNOSIS_CHAIN_API,
}

var RpcConfig = map[string]string{
	"mainnet": RPC_MAINNET,
	"goerli":  RPC_GOERLI,
	"xdai":    RPC_XDAI,
}

var SigningSchemeConfig = map[string]string{
	"mainnet": "eip712",
	"goerli":  "ethsign",
	"xdai":    "ethsign",
}

var Eip712OrderTypes = apitypes.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"Order": {
		{
			Name: "sellToken",
			Type: "address",
		},
		{
			Name: "buyToken",
			Type: "address",
		},
		{
			Name: "receiver",
			Type: "address",
		},
		{
			Name: "sellAmount",
			Type: "uint256",
		},
		{
			Name: "buyAmount",
			Type: "uint256",
		},
		{
			Name: "validTo",
			Type: "uint32",
		},
		{
			Name: "appData",
			Type: "bytes32",
		},
		{
			Name: "feeAmount",
			Type: "uint256",
		},
		{
			Name: "kind",
			Type: "string",
		},
		{
			Name: "partiallyFillable",
			Type: "bool",
		},
		{
			Name: "sellTokenBalance",
			Type: "string",
		},
		{
			Name: "buyTokenBalance",
			Type: "string",
		},
	},
}

// Default chainId is mainnet
var Domain = apitypes.TypedDataDomain{
	Name:              "Gnosis Protocol",
	Version:           "v2",
	ChainId:           math.NewHexOrDecimal256(1),
	VerifyingContract: GPv2Settlement,
}

var TypedData = apitypes.TypedData{
	Types:       Eip712OrderTypes,
	PrimaryType: "Order",
	Domain:      Domain,
	Message:     map[string]interface{}{},
}
