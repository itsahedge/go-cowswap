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
		CreationDate                 time.Time `json:"creationDate"`
		Owner                        string    `json:"owner"`
		UID                          string    `json:"uid"`
		AvailableBalance             string    `json:"availableBalance"`
		ExecutedBuyAmount            string    `json:"executedBuyAmount"`
		ExecutedSellAmount           string    `json:"executedSellAmount"`
		ExecutedSellAmountBeforeFees string    `json:"executedSellAmountBeforeFees"`
		ExecutedFeeAmount            string    `json:"executedFeeAmount"`
		Invalidated                  bool      `json:"invalidated"`
		Status                       string    `json:"status"`
		Class                        string    `json:"class"`
		SettlementContract           string    `json:"settlementContract"`
		FullFeeAmount                string    `json:"fullFeeAmount"`
		IsLiquidityOrder             bool      `json:"isLiquidityOrder"`
		SurplusFee                   string    `json:"surplusFee"`
		SurplusFeeTimestamp          time.Time `json:"surplusFeeTimestamp"`
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
		Interactions                 struct {
			Pre []interface{} `json:"pre"`
		} `json:"interactions"`
	} `json:"orders"`
	Prices struct {
		ZeroX6B175474E89094C44Da98B954Eedeac495271D0F string `json:"0x6b175474e89094c44da98b954eedeac495271d0f"`
		ZeroX853D955Acef822Db058Eb8505911Ed77F175B99E string `json:"0x853d955acef822db058eb8505911ed77f175b99e"`
		ZeroXa0B86991C6218B36C1D19D4A2E9Eb0Ce3606Eb48 string `json:"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"`
		ZeroXdac17F958D2Ee523A2206206994597C13D831Ec7 string `json:"0xdac17f958d2ee523a2206206994597c13d831ec7"`
		ZeroXdc8Af07A7861Bedd104B8093Ae3E9376Fc8596D2 string `json:"0xdc8af07a7861bedd104b8093ae3e9376fc8596d2"`
		ZeroXeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee string `json:"0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"`
		ZeroXfcc5C47Be19D06Bf83Eb04298B026F81069Ff65B string `json:"0xfcc5c47be19d06bf83eb04298b026f81069ff65b"`
	} `json:"prices"`
	Rewards struct {
		ZeroX20342797Ecfa9D84Fbf89A014A427A14F003Dea102Bae608Bd00651C6075Cf7Eb00098Ba6Eedaed1D4Ab31E7Fa14Cb969Ccce653637C4664 float64 `json:"0x20342797ecfa9d84fbf89a014a427a14f003dea102bae608bd00651c6075cf7eb00098ba6eedaed1d4ab31e7fa14cb969ccce653637c4664"`
		ZeroX6669Bddd136C833F575B2Dd1Cab130E0C9E96A3E84Cb17Dad00236C280B7C4Aa129436D21Ce486F18F4F008Fb6993E2F9Dbb6044637C4272 float64 `json:"0x6669bddd136c833f575b2dd1cab130e0c9e96a3e84cb17dad00236c280b7c4aa129436d21ce486f18f4f008fb6993e2f9dbb6044637c4272"`
		ZeroX6Cc70D125C55849264D5De893770C8A67F3Adfbed9Daba7E64879E2F4423Bc9Cd99087192Ab141F4Bef9Bc2937F3E6B16545A5Fd637C265D float64 `json:"0x6cc70d125c55849264d5de893770c8a67f3adfbed9daba7e64879e2f4423bc9cd99087192ab141f4bef9bc2937f3e6b16545a5fd637c265d"`
		ZeroXc7D1F4B64Aaf35Feec9B91F4F68092159Bb97D7D53197C47575817F631Ba83E1B0A9D7D6Db9D58Fd1Cf528353A746D4B126C8B13637C42Cf float64 `json:"0xc7d1f4b64aaf35feec9b91f4f68092159bb97d7d53197c47575817f631ba83e1b0a9d7d6db9d58fd1cf528353a746d4b126c8b13637c42cf"`
		ZeroXd3A1867Ed43Da50E2F0A89942B271B9415Bb4E1975D2Af7E4D7F009Cacae7Da17D2Ab9Ca511Ebd6F03971Fb417D3492Aa82513F06388E9Bd float64 `json:"0xd3a1867ed43da50e2f0a89942b271b9415bb4e1975d2af7e4d7f009cacae7da17d2ab9ca511ebd6f03971fb417d3492aa82513f06388e9bd"`
		ZeroXd9F0D9492Ab841Df1087C94Fa40028Fbd60Bc5818B2C1F802977D53D35659244937A0C4697E5551Ba3A50F8Beb279Ed0A71D08B5637C426F float64 `json:"0xd9f0d9492ab841df1087c94fa40028fbd60bc5818b2c1f802977d53d35659244937a0c4697e5551ba3a50f8beb279ed0a71d08b5637c426f"`
		ZeroXe8A8C82F8Bb05F139895Ac9A6Cb0Fdc37B087A360614796Be031D7B410Baecf4965A59359F0B59E03D622B1C384A71Aa18F77Af3637C2689 float64 `json:"0xe8a8c82f8bb05f139895ac9a6cb0fdc37b087a360614796be031d7b410baecf4965a59359f0b59e03d622b1c384a71aa18f77af3637c2689"`
		ZeroXe8D170Aa9661Eac38Fe8D0A952B5A01D956A1F22C933A47C419B09Ef2150027F6880F5334158980Ecfded17Ae18B455Efce1C0B5637C25Fb float64 `json:"0xe8d170aa9661eac38fe8d0a952b5a01d956a1f22c933a47c419b09ef2150027f6880f5334158980ecfded17ae18b455efce1c0b5637c25fb"`
	} `json:"rewards"`
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
