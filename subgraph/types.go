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
	Tokens []Token `json:"tokens"`
}

type Token struct {
	ID                  string `json:"id"`
	Address             string `json:"address"`
	FirstTradeTimestamp int    `json:"firstTradeTimestamp"`
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	Decimals            int    `json:"decimals"`
	TotalVolume         string `json:"totalVolume"`
	PriceEth            string `json:"priceEth"`
	PriceUsd            string `json:"priceUsd"`
	NumberOfTrades      int    `json:"numberOfTrades"`
	TotalVolumeUsd      string `json:"totalVolumeUsd"`
	TotalVolumeEth      string `json:"totalVolumeEth"`
}
