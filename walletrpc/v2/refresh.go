package walletrpc

import "context"

type RefreshRequest struct {
	// (Optional) The block height from which to start refreshing.
	StartHeight uint64 `json:"start_height,omitempty"`
}

type RefreshResponse struct {
	// Number of new blocks scanned.
	BlocksFetched uint64 `json:"blocks_fetched"`

	// States if transactions to the wallet have been found in the blocks.
	ReceivedMoney bool `json:"received_money"`
}

// Refresh a wallet after opening.
func (c *Client) Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error) {
	resp := &RefreshResponse{}
	err := c.Do(ctx, "refresh", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
