package walletrpc

type CheckTxKeyRequest struct {
	// Txid transaction id.
	Txid string `json:"txid"`

	// TxKey transaction secret key.
	TxKey string `json:"tx_key"`

	// Address destination public address of the transaction.
	Address string `json:"address"`
}

type CheckTxKeyResponse struct {
	// Confirmations Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`

	// InPool States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`

	// Received Amount of the transaction.
	Received uint64 `json:"received"`
}

// CheckTxKey Check a transaction in the blockchain with its secret key.
func (c *Client) CheckTxKey(req *CheckTxKeyRequest) (*CheckTxKeyResponse, error) {
	resp := &CheckTxKeyResponse{}
	err := c.Do("check_tx_key", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
