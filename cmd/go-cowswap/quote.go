package go_cowswap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

func (C *Client) Quote(o *QuoteReq) (*QuoteResponse, error) {
	bts, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/quote", C.Host), bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}

	resp, err := C.Http.Do(req)
	if err != nil {
		return nil, err
	}

	//return C.decodeQuoteResponse(resp)
	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200, 201:
		out := &QuoteResponse{}
		return out, dec.Decode(&out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
}
