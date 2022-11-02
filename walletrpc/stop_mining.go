package walletrpc

import "context"

// Stop mining in the Monero daemon.
func (c *Client) StopMining(ctx context.Context) error {
	return c.Do(ctx, "stop_mining", nil, nil)
}
