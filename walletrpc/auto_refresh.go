package walletrpc

type AutoRefreshRequest struct {
	// Enable (Optional) Enable or disable automatic refreshing (default = true).
	Enable bool `json:"enable"`

	// Period (Optional) The period of the wallet refresh cycle (i.e. time between refreshes) in seconds.
	Period uint64 `json:"period"`
}

// AutoRefresh Set whether and how often to automatically refresh the current wallet.
func (c *Client) AutoRefresh(req *AutoRefreshRequest) error {
	err := c.Do("auto_refresh", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
