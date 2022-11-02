package walletrpc

import "context"

type GetReserveProofRequest struct {
	// Proves all wallet balance to be disposable.
	All bool `json:"all"`

	// Specify the account from witch to prove reserve. (ignored if all is set to true)
	AccountIndex uint64 `json:"account_index"`

	// Amount (in atomic units) to prove the account has for reserve. (ignored if all is set to true)
	Amount uint64 `json:"amount"`

	// (Optional) Add a message to the signature to further authenticate the prooving process.
	Message string `json:"message,omitempty"`
}

type GetReserveProofResponse struct {
	// Reserve signature.
	Signature string `json:"signature"`
}

// Generate a signature to prove of an available amount in a wallet.
func (c *Client) GetReserveProof(ctx context.Context, req *GetReserveProofRequest) (*GetReserveProofResponse, error) {
	resp := &GetReserveProofResponse{}
	err := c.Do(ctx, "get_reserve_proof", &req, resp)
	return resp, err
}
