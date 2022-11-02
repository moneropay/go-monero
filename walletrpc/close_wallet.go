package walletrpc

import "context"

// Close the currently opened wallet, after trying to save it.
func (c *Client) CloseWallet(ctx context.Context) error {
	return c.Do(ctx, "close_wallet", nil, nil)
}
