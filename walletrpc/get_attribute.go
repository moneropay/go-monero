package walletrpc

type GetAttributeRequest struct {
	// Attribute name
	Key string `json:"key"`
}

type GetAttributeResponse struct {
	// Attribute value
	Value string `json:"value"`
}

// Get attribute value by name.
func (c *Client) GetAttribute(req *GetAttributeRequest) (*GetAttributeResponse, error) {
	resp := &GetAttributeResponse{}
	err := c.Do("get_attribute", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
