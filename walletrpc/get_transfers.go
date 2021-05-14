package walletrpc

type GetTransfersRequest struct {
	// In (defaults to false) Include incoming transfers.
	In bool `json:"in"`

	// Out (defaults to false) Include outgoing transfers.
	Out bool `json:"out"`

	// Pending (defaults to false) Include pending transfers.
	Pending bool `json:"pending"`

	// Failed (defaults to false) Include failed transfers.
	Failed bool `json:"failed"`

	// Pool (defaults to false) Include transfers from the daemon's transaction pool.
	Pool bool `json:"pool"`

	// FilterByHeight (Optional) Filter transfers by block height.
	FilterByHeight bool `json:"filter_by_height"`

	// MinHeight (Optional) Minimum block height to scan for transfers, if filtering by height is enabled.
	MinHeight uint64 `json:"min_height"`

	// MaxHeight (Optional) Maximum block height to scan for transfers, if filtering by height is enabled (defaults to max block height).
	MaxHeight uint64 `json:"max_height"`

	// AccountIndex (Optional) Index of the account to query for transfers. (defaults to 0)
	AccountIndex uint64 `json:"account_index"`

	// SubaddrIndices (Optional) List of subaddress indices to query for transfers. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices"`
}

type GetTransfersResponse struct {
	In      []Transfer `json:"in"`
	Out     []Transfer `json:"out"`
	Pending []Transfer `json:"pending"`
	Failed  []Transfer `json:"failed"`
	Pool    []Transfer `json:"pool"`
}

// GetTransfers Returns a list of transfers.
func (c *Client) GetTransfers(req *GetTransfersRequest) (*GetTransfersResponse, error) {
	resp := &GetTransfersResponse{}
	err := c.Do("get_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
