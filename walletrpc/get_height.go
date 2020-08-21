package walletrpc

type GetHeightResponse struct {
	// Height The current monero-wallet-rpc's blockchain height. If the wallet has been offline for a long time, it may need to catch up with the daemon.
	Height uint64 `json:"height"`
}

// GetHeight Returns the wallet's current block height.
func (c *Client) GetHeight() (*GetHeightResponse, error) {
	resp := &GetHeightResponse{}
	err := c.do("get_height", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
