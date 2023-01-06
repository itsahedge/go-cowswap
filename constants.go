package util

const (
	MAINNET_API        = "https://api.cow.fi/mainnet/api/v1"
	GOERLI_API         = "https://api.cow.fi/goerli/api/v1"
	GNOSIS_CHAIN_API   = "https://api.cow.fi/xdai/api/v1"
	GPv2Settlement     = "0x9008D19f58AAbD9eD0D60971565AA8510560ab41"
	GPv2_Vault_Relayer = "0xC92E8bdf79f0507f65a392b0ab4667716BFE0110"
)

const (
	RPC_MAINNET = "https://api.securerpc.com/v1"
	RPC_GOERLI  = "https://eth-goerli-rpc.gateway.pokt.network"
	RPC_XDAI    = "https://xdai-rpc.gateway.pokt.network"
)

const (
	SUBGRAPH_MAINNET      = "https://api.thegraph.com/subgraphs/name/cowprotocol/cow"
	SUBGRAPH_GOERLI       = "https://api.thegraph.com/subgraphs/name/cowprotocol/cow-goerli"
	SUBGRAPH_GNOSIS_CHAIN = "https://api.thegraph.com/subgraphs/name/cowprotocol/cow-gc"
)

var ETHEREUM_TOKEN_LIST = map[string]string{
	"WETH": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	"USDC": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	"COW":  "0xdef1ca1fb7fbcdc777520aa7f396b4e015f497ab",
	"GNO":  "0x6810e776880c02933d47db1b9fc05908e5386b96",
}

var GOERLI_TOKEN_LIST = map[string]string{
	"COW":  "0x91056D4A53E1faa1A84306D4deAEc71085394bC8",
	"GNO":  "0x02ABBDbAaa7b1BB64B5c878f7ac17f8DDa169532",
	"BAT":  "0x70cBa46d2e933030E2f274AE58c951C800548AeF",
	"DAI":  "0xdc31Ee1784292379Fbb2964b3B9C4124D8F89C60",
	"POLY": "0x9e32c0EfF886B6Ccae99350Fd5e7002dCED55F15",
	"UNI":  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
	"USDC": "0xD87Ba7A50B2E7E660f678A895E4B72E7CB4CCd9C",
	"WETH": "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6",
	"ZRX":  "0xe4E81Fa6B16327D4B78CFEB83AAdE04bA7075165",
}

var GNOSIS_SCAN_TOKEN_LIST = map[string]string{
	"GNO": "",
}
