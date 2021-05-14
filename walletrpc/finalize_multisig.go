package walletrpc

type FinalizeMultisigRequest struct {
	// MultisigInfo List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`

	// Password Wallet password
	Password string `json:"password"`
}

type FinalizeMultisigResponse struct {
	// Address multisig wallet address.
	Address string `json:"address"`
}

// FinalizeMultisig Turn this wallet into a multisig wallet, extra step for N-1/N wallets.
func (c *Client) FinalizeMultisig(req *FinalizeMultisigRequest) (*FinalizeMultisigResponse, error) {
	resp := &FinalizeMultisigResponse{}
	err := c.Do("finalize_multisig", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
