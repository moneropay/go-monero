package walletrpc

type RelayTxRequest struct {
	// Hex transaction metadata returned from a transfer method with get_tx_metadata set to true.
	Hex string `json:"hex"`
}

type RelayTxResponse struct {
	// TxHash for the publically searchable transaction hash.
	TxHash string `json:"tx_hash"`
}

// RelayTx Relay a transaction previously created with "do_not_relay":true.
func (c *Client) RelayTx(req *RelayTxRequest) (*RelayTxResponse, error) {
	resp := &RelayTxResponse{}
	err := c.do("relay_tx", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
