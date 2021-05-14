package walletrpc

type RefreshRequest struct {
	// StartHeight (Optional) The block height from which to start refreshing.
	StartHeight uint64 `json:"start_height"`
}

type RefreshResponse struct {
	// BlocksFetched Number of new blocks scanned.
	BlocksFetched uint64 `json:"blocks_fetched"`

	// ReceivedMoney States if transactions to the wallet have been found in the blocks.
	ReceivedMoney bool `json:"received_money"`
}

// Refresh Refresh a wallet after openning.
func (c *Client) Refresh(req *RefreshRequest) (*RefreshResponse, error) {
	resp := &RefreshResponse{}
	err := c.Do("refresh", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
