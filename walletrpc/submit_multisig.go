package walletrpc

import "context"

type SubmitMultisigRequest struct {
	// Multisig transaction in hex format, as returned by sign_multisig under tx_data_hex.
	TxDataHex string `json:"tx_data_hex"`
}

type SubmitMultisigResponse struct {
	// List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// Submit a signed multisig transaction.
func (c *Client) SubmitMultisig(ctx context.Context, req *SubmitMultisigRequest) (*SubmitMultisigResponse, error) {
	resp := &SubmitMultisigResponse{}
	err := c.Do(ctx, "submit_multisig", &req, resp)
	return resp, err
}
