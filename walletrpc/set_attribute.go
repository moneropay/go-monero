package walletrpc

type SetAttributeRequest struct {
	// Key attribute name
	Key string `json:"key"`

	// Value attribute value
	Value string `json:"value"`
}

// SetAttribute Set arbitrary attribute.
func (c *Client) SetAttribute(req *SetAttributeRequest) error {
	err := c.do("set_attribute", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
