package walletrpc

type CreateWalletRequest struct {
	// Wallet file name.
	Filename string `json:"filename"`

	// (Optional) password to protect the wallet.
	Password string `json:"password,omitempty"`

	// Language for your wallets' seed.
	Language string `json:"language"`
}

// Create a new wallet. You need to have set the argument "–wallet-dir" when launching monero-wallet-rpc to make this work.
func (c *Client) CreateWallet(req *CreateWalletRequest) error {
	err := c.Do("create_wallet", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
