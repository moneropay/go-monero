package walletrpc

type GetTransfersRequest struct {
	// Include incoming transfers. (Defaults to false)
	In bool `json:"in,omitempty"`

	// Include outgoing transfers. (Defaults to false)
	Out bool `json:"out,omitempty"`

	// Include pending transfers. (Defaults to false).
	Pending bool `json:"pending,omitempty"`

	// Include failed transfers. (Defaults to false)
	Failed bool `json:"failed,omitempty"`

	// Include transfers from the daemon's transaction pool. (Defaults to false)
	Pool bool `json:"pool,omitempty"`

	// (Optional) Filter transfers by block height.
	FilterByHeight bool `json:"filter_by_height,omitempty"`

	// (Optional) Minimum block height to scan for transfers, if filtering by height is enabled.
	MinHeight uint64 `json:"min_height,omitempty"`

	// (Optional) Maximum block height to scan for transfers, if filtering by height is enabled (Defaults to max block height).
	MaxHeight uint64 `json:"max_height,omitempty"`

	// (Optional) Index of the account to query for transfers. (Defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`

	// (Optional) List of subaddress indices to query for transfers. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
}

type GetTransfersResponse struct {
	In      []Transfer `json:"in"`
	Out     []Transfer `json:"out"`
	Pending []Transfer `json:"pending"`
	Failed  []Transfer `json:"failed"`
	Pool    []Transfer `json:"pool"`
}

// Returns a list of transfers.
func (c *Client) GetTransfers(req *GetTransfersRequest) (*GetTransfersResponse, error) {
	resp := &GetTransfersResponse{}
	err := c.Do("get_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
