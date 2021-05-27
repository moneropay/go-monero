package walletrpc

type GetSpendProofRequest struct {
	// Transaction id.
	Txid string `json:"txid"`

	// (Optional) Add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

type GetSpendProofResponse struct {
	// Spend signature.
	Signature string `json:"signature"`
}

// Generate a signature to prove a spend. Unlike proving a transaction, it does not requires the destination public address.
func (c *Client) GetSpendProof(req *GetSpendProofRequest) (*GetSpendProofResponse, error) {
	resp := &GetSpendProofResponse{}
	err := c.Do("get_spend_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
