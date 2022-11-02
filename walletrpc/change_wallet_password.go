package walletrpc

import "context"

type ChangeWalletPasswordRequest struct {
	// (Optional) Current wallet password, if defined.
	OldPassword string `json:"old_password,omitempty"`

	// (Optional) New wallet password, if not blank.
	NewPassword string `json:"new_password,omitempty"`
}

// Change a wallet password.
func (c *Client) ChangeWalletPassword(ctx context.Context, req *ChangeWalletPasswordRequest) error {
	return c.Do(ctx, "change_wallet_password", &req, nil)
}
