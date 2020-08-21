package walletrpc

type LabelAddressRequest struct {
	// Index containing the major & minor address index:
	Index SubaddressIndex `json:"index"`

	// Label  for the address.
	Label string `json:"label"`
}

// LabelAddress Label an address.
func (c *Client) LabelAddress(req *LabelAddressRequest) error {
	err := c.do("label_address", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
