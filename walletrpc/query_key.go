package walletrpc

import "context"

type QueryKeyRequest struct {
	// Which key to retrieve: "mnemonic" - the mnemonic seed (older wallets Do not have one) OR "view_key" - the view key
	KeyType string `json:"key_type"`
}

type QueryKeyResponse struct {
	// The view key will be hex encoded, while the mnemonic will be a string of words.
	Key string `json:"key"`
}

// Return the spend or view private key.
func (c *Client) QueryKey(ctx context.Context, req *QueryKeyRequest) (*QueryKeyResponse, error) {
	resp := &QueryKeyResponse{}
	err := c.Do(ctx, "query_key", &req, resp)
	return resp, err
}
