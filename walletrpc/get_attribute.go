package walletrpc

import "context"

type GetAttributeRequest struct {
	// Attribute name
	Key string `json:"key"`
}

type GetAttributeResponse struct {
	// Attribute value
	Value string `json:"value"`
}

// Get attribute value by name.
func (c *Client) GetAttribute(ctx context.Context, req *GetAttributeRequest) (*GetAttributeResponse, error) {
	resp := &GetAttributeResponse{}
	err := c.Do(ctx, "get_attribute", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
