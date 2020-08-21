package walletrpc

// Store Save the wallet file.
func (c *Client) Store() error {
	err := c.Do("store", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
