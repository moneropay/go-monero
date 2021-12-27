package walletrpc

import "context"

type GetAccountTagsResponse struct {
	// Array of account tag information:
	AccountTags []AccountTag `json:"account_tags"`
}

type AccountTag struct {
	// Filter tag.
	Tag string `json:"tag"`

	// Label for the tag.
	Label string `json:"label"`

	// List of tagged account indices.
	Accounts []int `json:"accounts"`
}

// Get a list of user-defined account tags.
func (c *Client) GetAccountTags(ctx context.Context) (*GetAccountTagsResponse, error) {
	resp := &GetAccountTagsResponse{}
	err := c.Do(ctx, "get_account_tags", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
