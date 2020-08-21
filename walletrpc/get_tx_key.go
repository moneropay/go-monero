package walletrpc

type GetTxKeyRequest struct {
	// Txid transaction id.
	Txid string `json:"txid"`
}

type GetTxKeyResponse struct {
	// TxKey transaction secret key.
	TxKey string `json:"tx_key"`
}

// GetTxKey Get transaction secret key from transaction id.
func (c *Client) GetTxKey(req *GetTxKeyRequest) (*GetTxKeyResponse, error) {
	resp := &GetTxKeyResponse{}
	err := c.do("get_tx_key", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
