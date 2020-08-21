package walletrpc

// StopMining Stop mining in the Monero daemon.
func (c *Client) StopMining() error {
	err := c.Do("stop_mining", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
