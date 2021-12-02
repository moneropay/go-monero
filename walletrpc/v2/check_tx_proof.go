package walletrpc

import "context"

type CheckTxProofRequest struct {
	// Transaction id.
	Txid string `json:"txid"`

	// Destination public address of the transaction.
	Address string `json:"address"`

	// (Optional) Should be the same message used in get_tx_proof.
	Message string `json:"message,omitempty"`

	// Transaction signature to confirm.
	Signature string `json:"signature"`
}

type CheckTxProofResponse struct {
	// Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`

	// States if the inputs proves the transaction.
	Good bool `json:"good"`

	// States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`

	// Amount of the transaction.
	Received uint64 `json:"received"`
}

// Prove a transaction by checking its signature.
func (c *Client) CheckTxProof(ctx context.Context, req *CheckTxProofRequest) (*CheckTxProofResponse, error) {
	resp := &CheckTxProofResponse{}
	err := c.Do(ctx, "check_tx_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
