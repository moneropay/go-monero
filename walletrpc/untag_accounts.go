package walletrpc

import "context"

type UntagAccountsRequest struct {
	// Remove tag from this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// Remove filtering tag from a list of accounts.
func (c *Client) UntagAccounts(ctx context.Context, req *UntagAccountsRequest) error {
	err := c.Do(ctx, "untag_accounts", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
