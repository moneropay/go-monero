package walletrpc

type GetTxProofRequest struct {
	// Txid transaction id.
	Txid string `json:"txid"`

	// Address destination public address of the transaction.
	Address string `json:"address"`

	// Message (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message"`
}

type GetTxProofResponse struct {
	// Signature transaction signature.
	Signature string `json:"signature"`
}

// GetTxProof Get transaction signature to prove it.
func (c *Client) GetTxProof(req *GetTxProofRequest) (*GetTxProofResponse, error) {
	resp := &GetTxProofResponse{}
	err := c.Do("get_tx_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
