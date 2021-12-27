package walletrpc

import "context"

type CreateAccountRequest struct {
	// (Optional) Label for the account.
	Label string `json:"label,omitempty"`
}

type CreateAccountResponse struct {
	// Index of the new account.
	AccountIndex uint64 `json:"account_index"`

	// Address for this account. Base58 representation of the public keys.
	Address string `json:"address"`
}

// Create a new account with an optional label.
func (c *Client) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := c.Do(ctx, "create_account", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
