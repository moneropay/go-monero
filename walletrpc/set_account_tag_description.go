package walletrpc

type SetAccountTagDescriptionRequest struct {
	// Tag Set a description for this tag.
	Tag string `json:"tag"`

	// Description Description for the tag.
	Description string `json:"description"`
}

// SetAccountTagDescription Set description for an account tag.
func (c *Client) SetAccountTagDescription(req *SetAccountTagDescriptionRequest) error {
	err := c.do("set_account_tag_description", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
