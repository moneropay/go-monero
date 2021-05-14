package walletrpc

type CheckSpendProofRequest struct {
	// Txid transaction id.
	Txid string `json:"txid"`

	// Message (Optional) Should be the same message used in get_spend_proof.
	Message string `json:"message"`

	// Signature spend signature to confirm.
	Signature string `json:"signature"`
}

type CheckSpendProofResponse struct {
	// Good States if the inputs proves the spend.
	Good bool `json:"good"`
}

// CheckSpendProof Prove a spend using a signature. Unlike proving a transaction, it does not requires the destination public address.
func (c *Client) CheckSpendProof(req *CheckSpendProofRequest) (*CheckSpendProofResponse, error) {
	resp := &CheckSpendProofResponse{}
	err := c.Do("check_spend_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
