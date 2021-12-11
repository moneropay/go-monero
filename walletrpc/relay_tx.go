package walletrpc

import "context"

type RelayTxRequest struct {
	// Transaction metadata returned from a transfer method with get_tx_metadata set to true.
	Hex string `json:"hex"`
}

type RelayTxResponse struct {
	// The publically searchable transaction hash.
	TxHash string `json:"tx_hash"`
}

// Relay a transaction previously created with "do_not_relay":true.
func (c *Client) RelayTx(ctx context.Context, req *RelayTxRequest) (*RelayTxResponse, error) {
	resp := &RelayTxResponse{}
	err := c.Do(ctx, "relay_tx", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
