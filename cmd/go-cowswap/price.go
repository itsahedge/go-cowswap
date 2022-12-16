package go_cowswap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NativePriceResponse struct {
	Price float64 `json:"price"`
}

func (C *Client) GetNativePrice(tokenAddress string) (*NativePriceResponse, error) {
	endpoint := fmt.Sprintf("%s/token/%s/native_price", C.Host, tokenAddress)
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
		out := &NativePriceResponse{}
		return out, dec.Decode(&out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
}
