package walletrpc

import "context"

type GetAccountsRequest struct {
	// (Optional) Tag for filtering accounts.
	Tag string `json:"tag,omitempty"`
}

type GetAccountsResponse struct {
	// Array of subaddress account information.
	SubaddressAccounts []Address `json:"subaddress_accounts"`

	// Total balance of the selected accounts (locked or unlocked).
	TotalBalance uint64 `json:"total_balance"`

	// Total unlocked balance of the selected accounts.
	TotalUnlockedBalance uint64 `json:"total_unlocked_balance"`
}

// Get all accounts for a wallet. Optionally filter accounts by tag.
func (c *Client) GetAccounts(ctx context.Context, req *GetAccountsRequest) (*GetAccountsResponse, error) {
	resp := &GetAccountsResponse{}
	err := c.Do(ctx, "get_accounts", &req, resp)
	return resp, err
}
