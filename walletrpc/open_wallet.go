package walletrpc

type OpenWalletRequest struct {
	// Filename wallet name stored in –wallet-dir.
	Filename string `json:"filename"`

	// Password (Optional) only needed if the wallet has a password defined.
	Password string `json:"password"`
}

// OpenWallet Open a wallet. You need to have set the argument "–wallet-dir" when launching monero-wallet-rpc to make this work.
func (c *Client) OpenWallet(req *OpenWalletRequest) error {
	err := c.do("open_wallet", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
