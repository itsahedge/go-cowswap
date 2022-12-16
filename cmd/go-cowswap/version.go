package go_cowswap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VersionResponse struct {
	Branch  string `json:"branch"`
	Commit  string `json:"commit"`
	Version string `json:"version"`
}

func (C *Client) Version() (*VersionResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/version", C.Host), nil)
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
		out := &VersionResponse{}
		return out, dec.Decode(out)
	default:
		err := &ErrorResponse{}
		if err2 := dec.Decode(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}
}
