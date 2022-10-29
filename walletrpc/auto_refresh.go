package walletrpc

import "context"

type AutoRefreshRequest struct {
	// (Optional) Enable or disable automatic refreshing (Defaults to true).
	Enable bool `json:"enable"`

	// (Optional) The period of the wallet refresh cycle (i.e. time between refreshes) in seconds.
	Period uint64 `json:"period,omitempty"`
}

// Set whether and how often to automatically refresh the current wallet.
func (c *Client) AutoRefresh(ctx context.Context, req *AutoRefreshRequest) error {
	return c.Do(ctx, "auto_refresh", &req, nil)
}
