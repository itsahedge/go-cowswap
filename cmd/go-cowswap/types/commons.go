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
