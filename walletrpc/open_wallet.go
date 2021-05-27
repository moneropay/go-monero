package walletrpc

type OpenWalletRequest struct {
	// Wallet name stored in –wallet-dir.
	Filename string `json:"filename"`

	// (Optional) Only needed if the wallet has a password defined.
	Password string `json:"password,omitempty"`
}

// Open a wallet. You need to have set the argument "–wallet-dir" when launching monero-wallet-rpc to make this work.
func (c *Client) OpenWallet(req *OpenWalletRequest) error {
	err := c.Do("open_wallet", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
