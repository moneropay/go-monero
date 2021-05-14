package walletrpc

type GetReserveProofRequest struct {
	// All Proves all wallet balance to be disposable.
	All bool `json:"all"`

	// AccountIndex Specify the account from witch to prove reserve. (ignored if all is set to true)
	AccountIndex uint64 `json:"account_index"`

	// Amount Amount (in atomic units) to prove the account has for reserve. (ignored if all is set to true)
	Amount uint64 `json:"amount"`

	// Message (Optional) add a message to the signature to further authenticate the prooving process.
	Message string `json:"message"`
}

type GetReserveProofResponse struct {
	// Signature reserve signature.
	Signature string `json:"signature"`
}

// GetReserveProof Generate a signature to prove of an available amount in a wallet.
func (c *Client) GetReserveProof(req *GetReserveProofRequest) (*GetReserveProofResponse, error) {
	resp := &GetReserveProofResponse{}
	err := c.Do("get_reserve_proof", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
