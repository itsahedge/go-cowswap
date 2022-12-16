package go_cowswap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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

func (C *Client) CreateOrder(o *CounterOrder) (*string, error) {
	bts, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s/orders", C.Host)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}

	resp, err := C.Http.Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	var out *string
	switch resp.StatusCode {
	case 200, 201:
		return out, dec.Decode(&out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
}
