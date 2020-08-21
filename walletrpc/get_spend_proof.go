package walletrpc

type GetSpendProofRequest struct {
	// Txid transaction id.
	Txid string `json:"txid"`

	// Message (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message"`
}

type GetSpendProofResponse struct {
	// Signature spend signature.
	Signature string `json:"signature"`
}

// GetSpendProof Generate a signature to prove a spend. Unlike proving a transaction, it does not requires the destination public address.
func (c *Client) GetSpendProof(req *GetSpendProofRequest) (*GetSpendProofResponse, error) {
	resp := &GetSpendProofResponse{}
	err := c.Do("get_spend_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
