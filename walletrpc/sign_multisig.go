package walletrpc

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
func (c *Client) SignMultisig(req *SignMultisigRequest) (*SignMultisigResponse, error) {
	resp := &SignMultisigResponse{}
	err := c.Do("sign_multisig", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
