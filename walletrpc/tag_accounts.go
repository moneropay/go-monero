package walletrpc

type TagAccountsRequest struct {
	// Tag for the accounts.
	Tag string `json:"tag"`

	// Tag this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// Apply a filtering tag to a list of accounts.
func (c *Client) TagAccounts(req *TagAccountsRequest) error {
	err := c.Do("tag_accounts", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
