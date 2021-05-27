package walletrpc

type GetTxKeyRequest struct {
	// Transaction id.
	Txid string `json:"txid"`
}

type GetTxKeyResponse struct {
	// Transaction secret key.
	TxKey string `json:"tx_key"`
}

// Get transaction secret key from transaction id.
func (c *Client) GetTxKey(req *GetTxKeyRequest) (*GetTxKeyResponse, error) {
	resp := &GetTxKeyResponse{}
	err := c.Do("get_tx_key", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
