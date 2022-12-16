package go_cowswap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

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

func (C *Client) GetOrdersByUid(uid string) (*OrderByUidResponse, error) {
	if uid == "" {
		return nil, errors.New("order UID not provided")
	}
	endpoint := fmt.Sprintf("%s/orders/%s", C.Host, uid)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := C.Http.Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200, 201:
		out := &OrderByUidResponse{}
		return out, dec.Decode(out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
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

func (C *Client) GetOrdersByTxHash(txHash string) (*OrdersByTxHashResponse, error) {
	if txHash == "" {
		return nil, errors.New("transaction hash not provided")
	}
	endpoint := fmt.Sprintf("%s/transactions/%s/orders", C.Host, txHash)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := C.Http.Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200, 201:
		out := &OrdersByTxHashResponse{}
		return out, dec.Decode(out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
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

func (C *Client) GetOrdersByUser(userAddress string, opts *OrdersPaginated) (*OrdersByTxHashResponse, error) {
	if userAddress == "" {
		return nil, errors.New("user address not provided")
	}
	endpoint := fmt.Sprintf("%s/account/%s/orders", C.Host, userAddress)
	fmt.Println("endpoint: ", endpoint)
	if opts != nil {
		if opts.Limit != "" && opts.Offset != "" {
			endpoint = fmt.Sprintf("%s?offset=%v&limit=%v", endpoint, opts.Offset, opts.Limit)
		} else {
			if opts.Limit != "" {
				endpoint = fmt.Sprintf("%s?limit=%v", endpoint, opts.Limit)
			}
			if opts.Offset != "" {
				endpoint = fmt.Sprintf("%s?offset=%v", endpoint, opts.Offset)
			}
		}
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("req:", req)

	resp, err := C.Http.Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200, 201:
		out := &OrdersByTxHashResponse{}
		return out, dec.Decode(out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
}
