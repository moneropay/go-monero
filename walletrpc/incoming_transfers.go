package walletrpc

type IncomingTransfersRequest struct {
	// "all": all the transfers, "available": only transfers which are not yet spent, OR "unavailable": only transfers which are already spent.
	TransferType string `json:"transfer_type"`

	// (Optional) Return transfers for this account. (defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`

	// (Optional) Return transfers sent to these subaddresses.
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
}

type IncomingTransfersResponse struct {
	// List of transfers
	Transfers []Transfer `json:"transfers"`
}

// Return a list of incoming transfers to the wallet.
func (c *Client) IncomingTransfers(req *IncomingTransfersRequest) (*IncomingTransfersResponse, error) {
	resp := &IncomingTransfersResponse{}
	err := c.Do("incoming_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
