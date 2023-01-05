package subgraph

type UserData struct {
	Users User `json:"user,omitempty"`
}

type Users struct {
	Users []User `json:"users,omitempty"`
}

type User struct {
	ID                  string         `json:"id,omitempty"`
	Address             string         `json:"address,omitempty"`
	FirstTradeTimestamp int64          `json:"firstTradeTimestamp,omitempty"`
	IsSolver            bool           `json:"isSolver,omitempty"`
	NumberOfTrades      int64          `json:"numberOfTrades,omitempty"`
	SolvedAmountEth     string         `json:"solvedAmountEth,omitempty"`
	SolvedAmountUsd     string         `json:"solvedAmountUsd,omitempty"`
	TradedAmountUsd     string         `json:"tradedAmountUsd,omitempty"`
	TradedAmountEth     string         `json:"tradedAmountEth,omitempty"`
	OrdersPlaced        []OrdersPlaced `json:"ordersPlaced,omitempty"`
}

type OrdersPlaced struct {
	ID       *string `json:"id,omitempty"`
	IsSigned *bool   `json:"isSigned,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens,omitempty"`
}

type Token struct {
	ID                  string `json:"id"`
	Address             string `json:"address,omitempty"`
	FirstTradeTimestamp int    `json:"firstTradeTimestamp,omitempty"`
	Name                string `json:"name,omitempty"`
	Symbol              string `json:"symbol,omitempty"`
	Decimals            int    `json:"decimals,omitempty"`
	TotalVolume         string `json:"totalVolume,omitempty"`
	PriceEth            string `json:"priceEth,omitempty"`
	PriceUsd            string `json:"priceUsd,omitempty"`
	NumberOfTrades      int    `json:"numberOfTrades,omitempty"`
	TotalVolumeUsd      string `json:"totalVolumeUsd,omitempty"`
	TotalVolumeEth      string `json:"totalVolumeEth,omitempty"`
}

type Orders struct {
	Orders []Order `json:"orders,omitempty"`
}

type Order struct {
	ID                  string  `json:"id"`
	TradesTimestamp     int     `json:"tradesTimestamp,omitempty"`
	InvalidateTimestamp int     `json:"invalidateTimestamp,omitempty"`
	PresignTimestamp    int     `json:"presignTimestamp,omitempty"`
	IsSigned            bool    `json:"isSigned,omitempty"`
	IsValid             bool    `json:"isValid,omitempty"`
	Owner               User    `json:"owner,omitempty"`
	Trades              []Trade `json:"trades,omitempty"`
}

type Trades struct {
	Trades []Trade `json:"trades,omitempty"`
}

type Trade struct {
	ID            string     `json:"id"`
	Timestamp     int        `json:"timestamp,omitempty"`
	GasPrice      string     `json:"gasPrice,omitempty"`
	FeeAmount     string     `json:"feeAmount,omitempty"`
	TxHash        string     `json:"txHash,omitempty"`
	Settlement    Settlement `json:"settlement,omitempty"`
	BuyAmount     string     `json:"buyAmount,omitempty"`
	SellAmount    string     `json:"sellAmount,omitempty"`
	BuyAmountUsd  string     `json:"buyAmountUsd,omitempty"`
	SellAmountUsd string     `json:"sellAmountUsd,omitempty"`
	BuyAmountEth  string     `json:"buyAmountEth,omitempty"`
	SellAmountEth string     `json:"sellAmountEth,omitempty"`
	SellToken     Token      `json:"sellToken,omitempty"`
	BuyToken      Token      `json:"buyToken,omitempty"`
	Order         Order      `json:"order,omitempty"`
}

type Settlement struct {
	ID                  string  `json:"id"`
	TxHash              string  `json:"txHash,omitempty"`
	FirstTradeTimestamp int     `json:"firstTradeTimestamp,omitempty"`
	Trades              []Trade `json:"trades,omitempty"`
	Solver              User    `json:"solver,omitempty"`
}

type Settlements struct {
	Settlements []Settlement `json:"settlements"`
}

type Bundle struct {
	ID          string `json:"id"`
	EthPriceUSD string `json:"ethPriceUSD"`
}

type Bundles struct {
	Bundles []Bundle `json:"bundles"`
}

type UniswapPools struct {
	UniswapPools []UniswapPool `json:"uniswapPools"`
}

type UniswapPool struct {
	ID                     string       `json:"id"`
	Liquidity              string       `json:"liquidity"`
	Token0Price            string       `json:"token0Price"`
	Token1Price            string       `json:"token1Price"`
	Tick                   string       `json:"tick"`
	TotalValueLockedToken0 string       `json:"totalValueLockedToken0"`
	TotalValueLockedToken1 string       `json:"totalValueLockedToken1"`
	Token0                 UniswapToken `json:"token0"`
	Token1                 UniswapToken `json:"token1"`
}

type UniswapToken struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	PriceEth string `json:"priceEth"`
	PriceUsd string `json:"priceUsd"`
}
