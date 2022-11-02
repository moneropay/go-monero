package walletrpc

import "context"

type GetTxKeyRequest struct {
	// Transaction id.
	Txid string `json:"txid"`
}

type GetTxKeyResponse struct {
	// Transaction secret key.
	TxKey string `json:"tx_key"`
}

// Get transaction secret key from transaction id.
func (c *Client) GetTxKey(ctx context.Context, req *GetTxKeyRequest) (*GetTxKeyResponse, error) {
	resp := &GetTxKeyResponse{}
	err := c.Do(ctx, "get_tx_key", &req, resp)
	return resp, err
}
