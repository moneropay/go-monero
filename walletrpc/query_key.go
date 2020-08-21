package walletrpc

type QueryKeyRequest struct {
	// KeyType Which key to retrieve: "mnemonic" - the mnemonic seed (older wallets do not have one) OR "view_key" - the view key
	KeyType string `json:"key_type"`
}

type QueryKeyResponse struct {
	// Key The view key will be hex encoded, while the mnemonic will be a string of words.
	Key string `json:"key"`
}

// QueryKey Return the spend or view private key.
func (c *Client) QueryKey(req *QueryKeyRequest) (*QueryKeyResponse, error) {
	resp := &QueryKeyResponse{}
	err := c.do("query_key", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
