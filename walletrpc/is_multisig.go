package walletrpc

type IsMultisigResponse struct {
	// States if the wallet is multisig
	Multisig bool `json:"multisig"`

	// Amount of signature needed to sign a transfer.
	Threshold uint64 `json:"threshold"`

	// Total amount of signature in the multisig wallet.
	Total uint64 `json:"total"`
}

// Check if a wallet is a multisig one.
func (c *Client) IsMultisig() (*IsMultisigResponse, error) {
	resp := &IsMultisigResponse{}
	err := c.Do("is_multisig", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
