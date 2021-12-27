package walletrpc

import "context"

// Close the currently opened wallet, after trying to save it.
func (c *Client) CloseWallet(ctx context.Context) error {
	err := c.Do(ctx, "close_wallet", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
