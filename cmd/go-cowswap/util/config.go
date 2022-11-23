package util

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/itsahedge/go-cowswap/cmd/go-cowswap/types"
)

const (
	MAINNET            = "https://api.cow.fi/mainnet/api/v1"
	GOERLI             = "https://api.cow.fi/goerli/api/v1"
	GNOSIS_CHAIN       = "https://api.cow.fi/xdai/api/v1"
	GPv2Settlement     = "0x9008D19f58AAbD9eD0D60971565AA8510560ab41"
	GPv2_Vault_Relayer = "0xC92E8bdf79f0507f65a392b0ab4667716BFE0110"
)

const (
	GOERLI_COW_TOKEN  = "0x91056D4A53E1faa1A84306D4deAEc71085394bC8"
	GOERLI_GNO_TOKEN  = "0x02ABBDbAaa7b1BB64B5c878f7ac17f8DDa169532"
	GOERLI_BAT_TOKEN  = "0x70cBa46d2e933030E2f274AE58c951C800548AeF"
	GOERLI_DAI_TOKEN  = "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60"
	GOERLI_POLY_TOKEN = "0x9e32c0EfF886B6Ccae99350Fd5e7002dCED55F15"
	GOERLI_UNI_TOKEN  = "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"
	GOERLI_USDC_TOKEN = "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C"
	GOERLI_WETH_TOKEN = "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"
	GOERLI_ZRX_TOKEN  = "0xe4E81Fa6B16327D4B78CFEB83AAdE04bA7075165"
)

const (
	WETH_TOKEN = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	COW_TOKEN  = "0xdef1ca1fb7fbcdc777520aa7f396b4e015f497ab"
	USDC_TOKEN = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	GNO_TOKEN  = "0x6810e776880c02933d47db1b9fc05908e5386b96"
)

var NetworkConfig = map[string]string{
	"mainnet": MAINNET,
	"goerli":  GOERLI,
	"xdai":    GNOSIS_CHAIN,
}

// Default options
var Options = types.Options{
	Network:    "mainnet",
	Host:       NetworkConfig["mainnet"],
	RpcUrl:     "https://rpc.flashbots.net",
	EthAddress: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
	PrivateKey: "",
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

var Domain = apitypes.TypedDataDomain{
	Name:              "Gnosis Protocol",
	Version:           "v2",
	ChainId:           math.NewHexOrDecimal256(1),
	VerifyingContract: GPv2Settlement,
}
