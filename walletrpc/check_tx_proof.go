package walletrpc

type CheckTxProofRequest struct {
	// Txid transaction id.
	Txid string `json:"txid"`

	// Address destination public address of the transaction.
	Address string `json:"address"`

	// Message (Optional) Should be the same message used in get_tx_proof.
	Message string `json:"message"`

	// Signature transaction signature to confirm.
	Signature string `json:"signature"`
}

type CheckTxProofResponse struct {
	// Confirmations Number of block mined after the one with the transaction.
	Confirmations uint64 `json:"confirmations"`

	// Good States if the inputs proves the transaction.
	Good bool `json:"good"`

	// InPool States if the transaction is still in pool or has been added to a block.
	InPool bool `json:"in_pool"`

	// Received Amount of the transaction.
	Received uint64 `json:"received"`
}

// CheckTxProof Prove a transaction by checking its signature.
func (c *Client) CheckTxProof(req *CheckTxProofRequest) (*CheckTxProofResponse, error) {
	resp := &CheckTxProofResponse{}
	err := c.Do("check_tx_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
