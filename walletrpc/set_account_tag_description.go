package walletrpc

import "context"

type SetAccountTagDescriptionRequest struct {
	// Set a description for this tag.
	Tag string `json:"tag"`

	// Description for the tag.
	Description string `json:"description"`
}

// Set description for an account tag.
func (c *Client) SetAccountTagDescription(ctx context.Context, req *SetAccountTagDescriptionRequest) error {
	return c.Do(ctx, "set_account_tag_description", &req, nil)
}
