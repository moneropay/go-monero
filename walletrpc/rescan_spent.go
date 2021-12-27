package walletrpc

import "context"

// Rescan the blockchain for spent outputs.
func (c *Client) RescanSpent(ctx context.Context) error {
	err := c.Do(ctx, "rescan_spent", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
