package walletrpc

import "context"

type GenerateFromKeysRequest struct {
	// (Optional) The block height to restore the wallet from. (Defaults to 0)
	RestoreHeight uint64 `json:"restore_height,omitempty"`

	// The wallet's file name on the RPC server.
	Filename string `json:"filename"`

	// The wallet's primary address.
	Address string `json:"address"`

	// (Optional) The wallet's private spend key. (Omit to create a view-only wallet)
	SpendKey string `json:"spendkey,omitempty"`

	// The wallet's private view key.
	ViewKey string `json:"viewkey"`

	// The wallet's password.
	Password string `json:"password"`

	// If true, save the current wallet before generating the new wallet. (Defaults to true)
	AutosaveCurrent bool `json:"autosave_current"`
}

type GenerateFromKeysResponse struct {
	// The wallet's address.
	Address string `json:"address"`

	// Verification message indicating that the wallet was generated successfully and whether or not it is a view-only wallet.
	Info string `json:"info"`
}

// Restores a wallet from a given wallet address, view key, and optional spend key.
func (c *Client) GenerateFromKeys(ctx context.Context, req *GenerateFromKeysRequest) (*GenerateFromKeysResponse, error) {
	resp := &GenerateFromKeysResponse{}
	err := c.Do(ctx, "generate_from_keys", &req, resp)
	return resp, err
}
