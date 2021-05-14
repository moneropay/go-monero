package walletrpc

// CloseWallet Close the currently opened wallet, after trying to save it.
func (c *Client) CloseWallet() error {
	err := c.Do("close_wallet", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
