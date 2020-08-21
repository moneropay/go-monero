package walletrpc

type UntagAccountsRequest struct {
	// Accounts Remove tag from this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// UntagAccounts Remove filtering tag from a list of accounts.
func (c *Client) UntagAccounts(req *UntagAccountsRequest) error {
	err := c.do("untag_accounts", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
