package walletrpc

type PrepareMultisigResponse struct {
	// MultisigInfo Multisig string to share with peers to create the multisig wallet.
	MultisigInfo string `json:"multisig_info"`
}

// PrepareMultisig Prepare a wallet for multisig by generating a multisig string to share with peers.
func (c *Client) PrepareMultisig() (*PrepareMultisigResponse, error) {
	resp := &PrepareMultisigResponse{}
	err := c.Do("prepare_multisig", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
