package walletrpc

type ChangeWalletPasswordRequest struct {
	// (Optional) Current wallet password, if defined.
	OldPassword string `json:"old_password,omitempty"`

	// (Optional) New wallet password, if not blank.
	NewPassword string `json:"new_password,omitempty"`
}

// Change a wallet password.
func (c *Client) ChangeWalletPassword(req *ChangeWalletPasswordRequest) error {
	err := c.Do("change_wallet_password", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
