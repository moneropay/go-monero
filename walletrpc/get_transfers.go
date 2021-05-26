package walletrpc

type GetTransfersRequest struct {
	// In (defaults to false) Include incoming transfers.
	In bool `json:"in,omitempty"`

	// Out (defaults to false) Include outgoing transfers.
	Out bool `json:"out,omitempty"`

	// Pending (defaults to false) Include pending transfers.
	Pending bool `json:"pending,omitempty"`

	// Failed (defaults to false) Include failed transfers.
	Failed bool `json:"failed,omitempty"`

	// Pool (defaults to false) Include transfers from the daemon's transaction pool.
	Pool bool `json:"pool,omitempty"`

	// FilterByHeight (Optional) Filter transfers by block height.
	FilterByHeight bool `json:"filter_by_height,omitempty"`

	// MinHeight (Optional) Minimum block height to scan for transfers, if filtering by height is enabled.
	MinHeight uint64 `json:"min_height,omitempty"`

	// MaxHeight (Optional) Maximum block height to scan for transfers, if filtering by height is enabled (defaults to max block height).
	MaxHeight uint64 `json:"max_height,omitempty"`

	// AccountIndex (Optional) Index of the account to query for transfers. (defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`

	// SubaddrIndices (Optional) List of subaddress indices to query for transfers. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
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
