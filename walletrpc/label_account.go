package walletrpc

type LabelAccountRequest struct {

	// AccountIndex Apply label to account at this index.
	AccountIndex uint64 `json:"account_index"`

	// Label Label for the account.
	Label string `json:"label"`
}

// LabelAccount Label an account.
func (c *Client) LabelAccount(req *LabelAccountRequest) error {
	err := c.Do("label_account", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
