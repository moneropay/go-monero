package walletrpc

import "context"

type SignMultisigRequest struct {
	// Multisig transaction in hex format, as returned by transfer under multisig_txset.
	TxDataHex string `json:"tx_data_hex"`
}

type SignMultisigResponse struct {
	// Multisig transaction in hex format.
	TxDataHex string `json:"tx_data_hex"`

	// List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// Sign a transaction in multisig.
func (c *Client) SignMultisig(ctx context.Context, req *SignMultisigRequest) (*SignMultisigResponse, error) {
	resp := &SignMultisigResponse{}
	err := c.Do(ctx, "sign_multisig", &req, resp)
	return resp, err
}
