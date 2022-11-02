package walletrpc

import "context"

type StartMiningRequest struct {
	// Number of threads created for mining.
	ThreadsCount uint64 `json:"threads_count"`

	// Allow to start the miner in smart mining mode.
	DoBackgroundMining bool `json:"do_background_mining"`

	// Ignore battery status (for smart mining only)
	IgnoreBattery bool `json:"ignore_battery"`
}

// Start mining in the Monero daemon.
func (c *Client) StartMining(ctx context.Context, req *StartMiningRequest) error {
	return c.Do(ctx, "start_mining", &req, nil)
}
