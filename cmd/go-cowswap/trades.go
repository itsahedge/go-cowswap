package go_cowswap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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

// GetTrades Exactly one of owner or order_uid has to be set.
func (C *Client) GetTrades(opts *GetTrades) (*TradesResponse, error) {
	if opts == nil {
		return nil, errors.New("must specify exactly one of owner or order_uid")
	}
	endpoint := fmt.Sprintf("%s/trades", C.Host)
	if opts != nil {
		if opts.Owner != "" && opts.OrderUid != "" {
			return nil, errors.New("must specify exactly one of owner or order_uid")
		}
		if opts.Owner != "" {
			endpoint = fmt.Sprintf("%s?owner=%s", endpoint, opts.Owner)
		}
		if opts.OrderUid != "" {
			endpoint = fmt.Sprintf("%s?orderUid=%s", endpoint, opts.OrderUid)
		}
	}
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
		out := &TradesResponse{}
		return out, dec.Decode(out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
}
