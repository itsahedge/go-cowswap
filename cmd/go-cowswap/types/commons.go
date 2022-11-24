package types

import (
	"time"
)

type Options struct {
	Network string
	Host    string
	RpcUrl  string

	EthAddress string
	PrivateKey string
}

type ApproveAllowance struct {
	TokenAddress  string `json:"token_address"`
	WalletAddress string `json:"wallet_address"`
}

type VersionResponse struct {
	Branch  string `json:"branch"`
	Commit  string `json:"commit"`
	Version string `json:"version"`
}

type QuoteReq struct {
	SellToken           string `json:"sellToken"`
	BuyToken            string `json:"buyToken"`
	Receiver            string `json:"receiver"`
	AppData             string `json:"appData"`
	PartiallyFillable   bool   `json:"partiallyFillable"`
	SellTokenBalance    string `json:"sellTokenBalance"`
	BuyTokenBalance     string `json:"buyTokenBalance"`
	PriceQuality        string `json:"priceQuality"`
	SigningScheme       string `json:"signingScheme"`
	OnchainOrder        bool   `json:"onchainOrder"`
	Kind                string `json:"kind"`
	SellAmountBeforeFee string `json:"sellAmountBeforeFee"`
	From                string `json:"from"`
}

type QuoteResponse struct {
	Quote struct {
		SellToken         string `json:"sellToken"`
		BuyToken          string `json:"buyToken"`
		Receiver          string `json:"receiver"`
		SellAmount        string `json:"sellAmount"`
		BuyAmount         string `json:"buyAmount"`
		ValidTo           int    `json:"validTo"`
		AppData           string `json:"appData"`
		FeeAmount         string `json:"feeAmount"`
		Kind              string `json:"kind"`
		PartiallyFillable bool   `json:"partiallyFillable"`
		SellTokenBalance  string `json:"sellTokenBalance"`
		BuyTokenBalance   string `json:"buyTokenBalance"`
		SigningScheme     string `json:"signingScheme"`
	} `json:"quote"`
	From       string    `json:"from"`
	Expiration time.Time `json:"expiration"`
	ID         int       `json:"id"`
}

type AuctionResponse struct {
	ID                    int `json:"id"`
	Block                 int `json:"block"`
	LatestSettlementBlock int `json:"latestSettlementBlock"`
	Orders                []struct {
		SellToken                    string    `json:"sellToken"`
		BuyToken                     string    `json:"buyToken"`
		Receiver                     string    `json:"receiver"`
		SellAmount                   string    `json:"sellAmount"`
		BuyAmount                    string    `json:"buyAmount"`
		ValidTo                      int       `json:"validTo"`
		AppData                      string    `json:"appData"`
		FeeAmount                    string    `json:"feeAmount"`
		Kind                         string    `json:"kind"`
		PartiallyFillable            bool      `json:"partiallyFillable"`
		SellTokenBalance             string    `json:"sellTokenBalance"`
		BuyTokenBalance              string    `json:"buyTokenBalance"`
		SigningScheme                string    `json:"signingScheme"`
		Signature                    string    `json:"signature"`
		From                         string    `json:"from"`
		QuoteID                      int       `json:"quoteId"`
		CreationTime                 time.Time `json:"creationTime"`
		Owner                        string    `json:"owner"`
		UID                          string    `json:"UID"`
		AvailableBalance             string    `json:"availableBalance"`
		ExecutedSellAmount           string    `json:"executedSellAmount"`
		ExecutedSellAmountBeforeFees string    `json:"executedSellAmountBeforeFees"`
		ExecutedBuyAmount            string    `json:"executedBuyAmount"`
		ExecutedFeeAmount            string    `json:"executedFeeAmount"`
		Invalidated                  bool      `json:"invalidated"`
		Status                       string    `json:"status"`
		FullFeeAmount                string    `json:"fullFeeAmount"`
		IsLiquidityOrder             bool      `json:"isLiquidityOrder"`
		EthflowData                  struct {
			IsRefunded  bool `json:"isRefunded"`
			UserValidTo int  `json:"userValidTo"`
		} `json:"ethflowData"`
		OnchainUser string `json:"onchainUser"`
	} `json:"orders"`
	Prices struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"prices"`
}

type NativePriceResponse struct {
	Price float64 `json:"price"`
}

type GetTrades struct {
	Owner    string
	OrderUid string
}

type TradesResponse []struct {
	BlockNumber          int    `json:"blockNumber"`
	LogIndex             int    `json:"logIndex"`
	OrderUID             string `json:"orderUid"`
	BuyAmount            string `json:"buyAmount"`
	SellAmount           string `json:"sellAmount"`
	SellAmountBeforeFees string `json:"sellAmountBeforeFees"`
	Owner                string `json:"owner"`
	BuyToken             string `json:"buyToken"`
	SellToken            string `json:"sellToken"`
	TxHash               string `json:"txHash"`
}

type SolverCompetitionResponse struct {
	AuctionId                  int     `json:"auctionId"`
	TransactionHash            string  `json:"transactionHash"`
	GasPrice                   float64 `json:"gasPrice"`
	LiquidityCollectedBlock    int     `json:"liquidityCollectedBlock"`
	CompetitionSimulationBlock int     `json:"competitionSimulationBlock"`
	Solutions                  []struct {
		Solver    string `json:"solver"`
		Objective struct {
			Total   float64 `json:"total"`
			Surplus float64 `json:"surplus"`
			Fees    float64 `json:"fees"`
			Cost    float64 `json:"cost"`
			Gas     int     `json:"gas"`
		} `json:"objective"`
		Prices struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"prices"`
		Orders []struct {
			Id             string `json:"id"`
			ExecutedAmount string `json:"executedAmount"`
		} `json:"orders"`
		CallData string `json:"callData"`
	} `json:"solutions"`
}

type OrderByUidResponse struct {
	SellToken                    string    `json:"sellToken"`
	BuyToken                     string    `json:"buyToken"`
	Receiver                     string    `json:"receiver"`
	SellAmount                   string    `json:"sellAmount"`
	BuyAmount                    string    `json:"buyAmount"`
	ValidTo                      int       `json:"validTo"`
	AppData                      string    `json:"appData"`
	FeeAmount                    string    `json:"feeAmount"`
	Kind                         string    `json:"kind"`
	PartiallyFillable            bool      `json:"partiallyFillable"`
	SellTokenBalance             string    `json:"sellTokenBalance"`
	BuyTokenBalance              string    `json:"buyTokenBalance"`
	SigningScheme                string    `json:"signingScheme"`
	Signature                    string    `json:"signature"`
	From                         string    `json:"from"`
	QuoteID                      int       `json:"quoteId"`
	CreationTime                 time.Time `json:"creationTime"`
	Owner                        string    `json:"owner"`
	UID                          string    `json:"UID"`
	AvailableBalance             string    `json:"availableBalance"`
	ExecutedSellAmount           string    `json:"executedSellAmount"`
	ExecutedSellAmountBeforeFees string    `json:"executedSellAmountBeforeFees"`
	ExecutedBuyAmount            string    `json:"executedBuyAmount"`
	ExecutedFeeAmount            string    `json:"executedFeeAmount"`
	Invalidated                  bool      `json:"invalidated"`
	Status                       string    `json:"status"`
	FullFeeAmount                string    `json:"fullFeeAmount"`
	IsLiquidityOrder             bool      `json:"isLiquidityOrder"`
	EthflowData                  struct {
		IsRefunded  bool `json:"isRefunded"`
		UserValidTo int  `json:"userValidTo"`
	} `json:"ethflowData"`
	OnchainUser string `json:"onchainUser"`
}

type OrdersByTxHashResponse []struct {
	CreationDate                 time.Time   `json:"creationDate"`
	Owner                        string      `json:"owner"`
	UID                          string      `json:"uid"`
	AvailableBalance             interface{} `json:"availableBalance"`
	ExecutedBuyAmount            string      `json:"executedBuyAmount"`
	ExecutedSellAmount           string      `json:"executedSellAmount"`
	ExecutedSellAmountBeforeFees string      `json:"executedSellAmountBeforeFees"`
	ExecutedFeeAmount            string      `json:"executedFeeAmount"`
	Invalidated                  bool        `json:"invalidated"`
	Status                       string      `json:"status"`
	Class                        string      `json:"class"`
	SettlementContract           string      `json:"settlementContract"`
	FullFeeAmount                string      `json:"fullFeeAmount"`
	IsLiquidityOrder             bool        `json:"isLiquidityOrder"`
	SellToken                    string      `json:"sellToken"`
	BuyToken                     string      `json:"buyToken"`
	Receiver                     string      `json:"receiver"`
	SellAmount                   string      `json:"sellAmount"`
	BuyAmount                    string      `json:"buyAmount"`
	ValidTo                      int         `json:"validTo"`
	AppData                      string      `json:"appData"`
	FeeAmount                    string      `json:"feeAmount"`
	Kind                         string      `json:"kind"`
	PartiallyFillable            bool        `json:"partiallyFillable"`
	SellTokenBalance             string      `json:"sellTokenBalance"`
	BuyTokenBalance              string      `json:"buyTokenBalance"`
	SigningScheme                string      `json:"signingScheme"`
	Signature                    string      `json:"signature"`
	Interactions                 struct {
		Pre []interface{} `json:"pre"`
	} `json:"interactions"`
}

type OrdersPaginated struct {
	Offset string
	Limit  string
}

type OrdersByUserResponse []struct {
	CreationDate                 time.Time   `json:"creationDate"`
	Owner                        string      `json:"owner"`
	UID                          string      `json:"uid"`
	AvailableBalance             interface{} `json:"availableBalance"`
	ExecutedBuyAmount            string      `json:"executedBuyAmount"`
	ExecutedSellAmount           string      `json:"executedSellAmount"`
	ExecutedSellAmountBeforeFees string      `json:"executedSellAmountBeforeFees"`
	ExecutedFeeAmount            string      `json:"executedFeeAmount"`
	Invalidated                  bool        `json:"invalidated"`
	Status                       string      `json:"status"`
	Class                        string      `json:"class"`
	SettlementContract           string      `json:"settlementContract"`
	FullFeeAmount                string      `json:"fullFeeAmount"`
	IsLiquidityOrder             bool        `json:"isLiquidityOrder"`
	SellToken                    string      `json:"sellToken"`
	BuyToken                     string      `json:"buyToken"`
	Receiver                     string      `json:"receiver"`
	SellAmount                   string      `json:"sellAmount"`
	BuyAmount                    string      `json:"buyAmount"`
	ValidTo                      int         `json:"validTo"`
	AppData                      string      `json:"appData"`
	FeeAmount                    string      `json:"feeAmount"`
	Kind                         string      `json:"kind"`
	PartiallyFillable            bool        `json:"partiallyFillable"`
	SellTokenBalance             string      `json:"sellTokenBalance"`
	BuyTokenBalance              string      `json:"buyTokenBalance"`
	SigningScheme                string      `json:"signingScheme"`
	Signature                    string      `json:"signature"`
	Interactions                 struct {
		Pre []interface{} `json:"pre"`
	} `json:"interactions"`
}

type SolverAuctionByTxHashResponse struct {
	AuctionId                  int     `json:"auctionId"`
	TransactionHash            string  `json:"transactionHash"`
	GasPrice                   float64 `json:"gasPrice"`
	LiquidityCollectedBlock    int     `json:"liquidityCollectedBlock"`
	CompetitionSimulationBlock int     `json:"competitionSimulationBlock"`
	Solutions                  []struct {
		Solver    string `json:"solver"`
		Objective struct {
			Total   float64 `json:"total"`
			Surplus float64 `json:"surplus"`
			Fees    float64 `json:"fees"`
			Cost    float64 `json:"cost"`
			Gas     int     `json:"gas"`
		} `json:"objective"`
		Prices struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"prices"`
		Orders []struct {
			Id             string `json:"id"`
			ExecutedAmount string `json:"executedAmount"`
		} `json:"orders"`
		CallData string `json:"callData"`
	} `json:"solutions"`
}

// CounterOrder represents a Gnosis CounterOrder.
type CounterOrder struct {
	SellToken         string `json:"sellToken,omitempty"`
	BuyToken          string `json:"buyToken,omitempty"`
	Receiver          string `json:"receiver,omitempty"`
	SellAmount        string `json:"sellAmount"`
	BuyAmount         string `json:"buyAmount"`
	ValidTo           uint32 `json:"validTo,omitempty"`
	AppData           string `json:"appData,omitempty"`
	FeeAmount         string `json:"feeAmount"`
	Kind              string `json:"kind,omitempty"`
	PartiallyFillable bool   `json:"partiallyFillable"`
	Signature         string `json:"signature,omitempty"`
	SigningScheme     string `json:"signingScheme,omitempty"`
	SellTokenBalance  string `json:"sellTokenBalance,omitempty"`
	BuyTokenBalance   string `json:"buyTokenBalance,omitempty"`
	From              string `json:"from,omitempty"`
}
