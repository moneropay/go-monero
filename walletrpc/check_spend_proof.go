package walletrpc

import "context"

type CheckSpendProofRequest struct {
	// Transaction id.
	Txid string `json:"txid"`

	// (Optional) Should be the same message used in get_spend_proof.
	Message string `json:"message,omitempty"`

	// Spend signature to confirm.
	Signature string `json:"signature"`
}

type CheckSpendProofResponse struct {
	// States if the inputs proves the spend.
	Good bool `json:"good"`
}

// Prove a spend using a signature. Unlike proving a transaction, it does not requires the destination public address.
func (c *Client) CheckSpendProof(ctx context.Context, req *CheckSpendProofRequest) (*CheckSpendProofResponse, error) {
	resp := &CheckSpendProofResponse{}
	err := c.Do(ctx, "check_spend_proof", &req, resp)
	return resp, err
}
