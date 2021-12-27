package walletrpc

import "context"

type MakeMultisigRequest struct {
	// List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`

	// Amount of signatures needed to sign a transfer. Must be less or equal than the amount of signature in multisig_info.
	Threshold uint64 `json:"threshold"`

	// Wallet password
	Password string `json:"password"`
}

type MakeMultisigResponse struct {
	// Multisig wallet address.
	Address string `json:"address"`

	// Multisig string to share with peers to create the multisig wallet (extra step for N-1/N wallets).
	MultisigInfo string `json:"multisig_info"`
}

// Make a wallet multisig by importing peers multisig string.
func (c *Client) MakeMultisig(ctx context.Context, req *MakeMultisigRequest) (*MakeMultisigResponse, error) {
	resp := &MakeMultisigResponse{}
	err := c.Do(ctx, "make_multisig", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
