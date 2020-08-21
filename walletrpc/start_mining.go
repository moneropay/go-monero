package walletrpc

type StartMiningRequest struct {
	// ThreadsCount Number of threads created for mining.
	ThreadsCount uint64 `json:"threads_count"`

	// DoBackgroundMining Allow to start the miner in smart mining mode.
	DoBackgroundMining bool `json:"do_background_mining"`

	// IgnoreBattery Ignore battery status (for smart mining only)
	IgnoreBattery bool `json:"ignore_battery"`
}

// StartMining Start mining in the Monero daemon.
func (c *Client) StartMining(req *StartMiningRequest) error {
	err := c.do("start_mining", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
