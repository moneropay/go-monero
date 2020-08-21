package walletrpc

type ChangeWalletPasswordRequest struct {
	// OldPassword (Optional) Current wallet password, if defined.
	OldPassword string `json:"old_password"`

	// NewPassword (Optional) New wallet password, if not blank.
	NewPassword string `json:"new_password"`
}

// ChangeWalletPassword Change a wallet password.
func (c *Client) ChangeWalletPassword(req *ChangeWalletPasswordRequest) error {
	err := c.do("change_wallet_password", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
