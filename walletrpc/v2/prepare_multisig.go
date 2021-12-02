package walletrpc

import "context"

type PrepareMultisigResponse struct {
	// Multisig string to share with peers to create the multisig wallet.
	MultisigInfo string `json:"multisig_info"`
}

// Prepare a wallet for multisig by generating a multisig string to share with peers.
func (c *Client) PrepareMultisig(ctx context.Context) (*PrepareMultisigResponse, error) {
	resp := &PrepareMultisigResponse{}
	err := c.Do(ctx, "prepare_multisig", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
