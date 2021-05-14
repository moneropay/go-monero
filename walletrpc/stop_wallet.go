package walletrpc

// StopWallet Stops the wallet, storing the current state.
func (c *Client) StopWallet() error {
	err := c.Do("stop_wallet", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
