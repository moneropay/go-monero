package walletrpc

type SignMultisigRequest struct {
	// TxDataHex Multisig transaction in hex format, as returned by transfer under multisig_txset.
	TxDataHex string `json:"tx_data_hex"`
}

type SignMultisigResponse struct {
	// TxDataHex Multisig transaction in hex format.
	TxDataHex string `json:"tx_data_hex"`

	// TxHashList List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// SignMultisig Sign a transaction in multisig.
func (c *Client) SignMultisig(req *SignMultisigRequest) (*SignMultisigResponse, error) {
	resp := &SignMultisigResponse{}
	err := c.Do("sign_multisig", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
