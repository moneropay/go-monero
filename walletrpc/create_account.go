package walletrpc

type CreateAccountRequest struct {
	// Label (Optional) for the account.
	Label string `json:"label"`
}

type CreateAccountResponse struct {
	// AccountIndex Index of the new account.
	AccountIndex uint64 `json:"account_index"`

	// Address  for this account. Base58 representation of the public keys.
	Address string `json:"address"`
}

// CreateAccount Create a new account with an optional label.
func (c *Client) CreateAccount(req *CreateAccountRequest) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := c.do("create_account", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
