package walletrpc

type SetAttributeRequest struct {
	// Attribute name
	Key string `json:"key"`

	// Attribute value
	Value string `json:"value"`
}

// Set arbitrary attribute.
func (c *Client) SetAttribute(req *SetAttributeRequest) error {
	err := c.Do("set_attribute", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
