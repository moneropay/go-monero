package walletrpc

type MakeMultisigRequest struct {
	// MultisigInfo List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`

	// Threshold Amount of signatures needed to sign a transfer. Must be less or equal than the amount of signature in multisig_info.
	Threshold uint64 `json:"threshold"`

	// Password Wallet password
	Password string `json:"password"`
}

type MakeMultisigResponse struct {
	// Address multisig wallet address.
	Address string `json:"address"`

	// MultisigInfo Multisig string to share with peers to create the multisig wallet (extra step for N-1/N wallets).
	MultisigInfo string `json:"multisig_info"`
}

// MakeMultisig Make a wallet multisig by importing peers multisig string.
func (c *Client) MakeMultisig(req *MakeMultisigRequest) (*MakeMultisigResponse, error) {
	resp := &MakeMultisigResponse{}
	err := c.do("make_multisig", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
