package walletrpc

type GetAccountsRequest struct {
	// Tag (Optional) Tag for filtering accounts.
	Tag string `json:"tag"`
}

type GetAccountsResponse struct {
	// SubaddressAccounts information:
	SubaddressAccounts []Address `json:"subaddress_accounts"`

	// TotalBalance Total balance of the selected accounts (locked or unlocked).
	TotalBalance uint64 `json:"total_balance"`

	// TotalUnlockedBalance Total unlocked balance of the selected accounts.
	TotalUnlockedBalance uint64 `json:"total_unlocked_balance"`
}

// GetAccounts Get all accounts for a wallet. Optionally filter accounts by tag.
func (c *Client) GetAccounts(req *GetAccountsRequest) (*GetAccountsResponse, error) {
	resp := &GetAccountsResponse{}
	err := c.do("get_accounts", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
