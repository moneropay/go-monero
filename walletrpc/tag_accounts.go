package walletrpc

type TagAccountsRequest struct {
	// Tag for the accounts.
	Tag string `json:"tag"`

	// Accounts Tag this list of accounts.
	Accounts []uint64 `json:"accounts"`
}

// TagAccounts Apply a filtering tag to a list of accounts.
func (c *Client) TagAccounts(req *TagAccountsRequest) error {
	err := c.do("tag_accounts", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
