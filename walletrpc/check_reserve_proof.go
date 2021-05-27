package walletrpc

type CheckReserveProofRequest struct {
	// Public address of the wallet.
	Address string `json:"address"`

	// (Optional) Should be the same message used in get_reserve_proof.
	Message string `json:"message,omitempty"`

	// Reserve signature to confirm.
	Signature string `json:"signature"`
}

type CheckReserveProofResponse struct {
	// States if the inputs proves the reserve.
	Good bool `json:"good"`
}

// Proves a wallet has a disposable reserve using a signature.
func (c *Client) CheckReserveProof(req *CheckReserveProofRequest) (*CheckReserveProofResponse, error) {
	resp := &CheckReserveProofResponse{}
	err := c.Do("check_reserve_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
