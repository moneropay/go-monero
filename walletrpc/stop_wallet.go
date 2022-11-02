package walletrpc

import "context"

// Stops the wallet, storing the current state.
func (c *Client) StopWallet(ctx context.Context) error {
	return c.Do(ctx, "stop_wallet", nil, nil)
}
