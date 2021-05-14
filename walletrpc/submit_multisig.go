package walletrpc

type SubmitMultisigRequest struct {
	// TxDataHex Multisig transaction in hex format, as returned by sign_multisig under tx_data_hex.
	TxDataHex string `json:"tx_data_hex"`
}

type SubmitMultisigResponse struct {
	// TxHashList List of transaction Hash.
	TxHashList []string `json:"tx_hash_list"`
}

// SubmitMultisig Submit a signed multisig transaction.
func (c *Client) SubmitMultisig(req *SubmitMultisigRequest) (*SubmitMultisigResponse, error) {
	resp := &SubmitMultisigResponse{}
	err := c.Do("submit_multisig", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
