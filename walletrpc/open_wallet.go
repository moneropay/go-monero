package walletrpc

import "context"

type OpenWalletRequest struct {
	// Wallet name stored in –wallet-dir.
	Filename string `json:"filename"`

	// (Optional) Only needed if the wallet has a password defined.
	Password string `json:"password,omitempty"`
}

// Open a wallet. You need to have set the argument "–wallet-dir" when launching monero-wallet-rpc to make this work.
func (c *Client) OpenWallet(ctx context.Context, req *OpenWalletRequest) error {
	return c.Do(ctx, "open_wallet", &req, nil)
}
