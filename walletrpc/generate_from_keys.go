package walletrpc

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

// Restores a wallet from a given wallet address, view key, and optional spend key.
func (c *Client) CreateWallet(req *CreateWalletRequest) error {
	err := c.Do("create_wallet", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
