package walletrpc

import "context"

type SetAttributeRequest struct {
	// Attribute name
	Key string `json:"key"`

	// Attribute value
	Value string `json:"value"`
}

// Set arbitrary attribute.
func (c *Client) SetAttribute(ctx context.Context, req *SetAttributeRequest) error {
	return c.Do(ctx, "set_attribute", &req, nil)
}
