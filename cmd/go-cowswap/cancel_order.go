package go_cowswap

type CancelOrder struct {
	OrderUidsStr  string `json:"order_uids_string"`
	OrderUids     []byte `json:"order_uids"`
	Signature     string `json:"signature"`
	SigningScheme string `json:"signing_scheme"`
}

//TODO add http handler
