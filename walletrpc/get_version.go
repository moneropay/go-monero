package walletrpc

type GetVersionResponse struct {
	// Version RPC version, formatted with Major * 2^16 + Minor (Major encoded over the first 16 bits, and Minor over the last 16 bits).
	Version uint64 `json:"version"`
}

// GetVersion Get RPC version Major & Minor integer-format, where Major is the first 16 bits and Minor the last 16 bits.
func (c *Client) GetVersion() (*GetVersionResponse, error) {
	resp := &GetVersionResponse{}
	err := c.do("get_version", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
