package walletrpc

import "context"

// Save the wallet file.
func (c *Client) Store(ctx context.Context) error {
	err := c.Do(ctx, "store", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
