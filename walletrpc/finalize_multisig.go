package walletrpc

import "context"

type FinalizeMultisigRequest struct {
	// List of multisig string from peers.
	MultisigInfo []string `json:"multisig_info"`

	// Wallet password
	Password string `json:"password"`
}

type FinalizeMultisigResponse struct {
	// Multisig wallet address.
	Address string `json:"address"`
}

// Turn this wallet into a multisig wallet, extra step for N-1/N wallets.
func (c *Client) FinalizeMultisig(ctx context.Context, req *FinalizeMultisigRequest) (*FinalizeMultisigResponse, error) {
	resp := &FinalizeMultisigResponse{}
	err := c.Do(ctx, "finalize_multisig", &req, resp)
	return resp, err
}
