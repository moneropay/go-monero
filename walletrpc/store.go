package walletrpc

// Store Save the wallet file.
func (c *Client) Store() error {
	err := c.do("store", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
