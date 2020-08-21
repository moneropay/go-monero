package walletrpc

type GetAccountTagsResponse struct {
	// AccountTags array of account tag information:
	AccountTags []AccountTag `json:"account_tags"`
}

type AccountTag struct {
	// Tag Filter tag.
	Tag string `json:"tag"`

	// Label Label for the tag.
	Label string `json:"label"`

	// Accounts List of tagged account indices.
	Accounts []int `json:"accounts"`
}

// GetAccountTags Get a list of user-defined account tags.
func (c *Client) GetAccountTags() (*GetAccountTagsResponse, error) {
	resp := &GetAccountTagsResponse{}
	err := c.do("get_account_tags", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
