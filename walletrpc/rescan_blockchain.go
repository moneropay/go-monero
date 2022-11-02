package walletrpc

import "context"

// Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself.
func (c *Client) RescanBlockchain(ctx context.Context) error {
	return c.Do(ctx, "rescan_blockchain", nil, nil)
}
