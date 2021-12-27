package walletrpc

import "context"

type CheckTxKeyRequest struct {
	// Transaction id.
	Txid string `json:"txid"`

	// Transaction secret key.
	TxKey string `json:"tx_key"`

	// Destination public address of the transaction.
	Address string `json:"address"`
}

type CheckTxKeyResponse struct {
	// Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`

	// States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`

	// Amount of the transaction.
	Received uint64 `json:"received"`
}

// Check a transaction in the blockchain with its secret key.
func (c *Client) CheckTxKey(ctx context.Context, req *CheckTxKeyRequest) (*CheckTxKeyResponse, error) {
	resp := &CheckTxKeyResponse{}
	err := c.Do(ctx, "check_tx_key", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
