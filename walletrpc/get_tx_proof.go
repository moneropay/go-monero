package walletrpc

import "context"

type GetTxProofRequest struct {
	// Transaction id.
	Txid string `json:"txid"`

	// Destination public address of the transaction.
	Address string `json:"address"`

	// (Optional) add a message to the signature to further authenticate the proving process.
	Message string `json:"message,omitempty"`
}

type GetTxProofResponse struct {
	// Transaction signature.
	Signature string `json:"signature"`
}

// Get transaction signature to prove it.
func (c *Client) GetTxProof(ctx context.Context, req *GetTxProofRequest) (*GetTxProofResponse, error) {
	resp := &GetTxProofResponse{}
	err := c.Do(ctx, "get_tx_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
