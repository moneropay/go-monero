package walletrpc

import "context"

// Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself.
func (c *Client) RescanBlockchain(ctx context.Context) error {
	err := c.Do(ctx, "rescan_blockchain", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
