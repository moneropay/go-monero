package walletrpc

import "context"

type TagAccountsRequest struct {
	// Tag for the accounts.
	Tag string `json:"tag"`

	// Tag this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// Apply a filtering tag to a list of accounts.
func (c *Client) TagAccounts(ctx context.Context, req *TagAccountsRequest) error {
	return c.Do(ctx, "tag_accounts", &req, nil)
}
