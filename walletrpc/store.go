package walletrpc

import "context"

// Save the wallet file.
func (c *Client) Store(ctx context.Context) error {
	return c.Do(ctx, "store", nil, nil)
}
