package walletrpc

type IncomingTransfersRequest struct {
	// TransferType "all": all the transfers, "available": only transfers which are not yet spent, OR "unavailable": only transfers which are already spent.
	TransferType string `json:"transfer_type"`

	// AccountIndex (Optional) Return transfers for this account. (defaults to 0)
	AccountIndex uint64 `json:"account_index"`

	// SubaddrIndices (Optional) Return transfers sent to these subaddresses.
	SubaddrIndices []uint64 `json:"subaddr_indices"`

	// Verbose (Optional) Enable verbose output, return key image if true.
	Verbose bool `json:"verbose"`
}

type IncomingTransfersResponse struct {
	// Transfers list of transfers
	Transfers []Transfer `json:"transfers"`
}

// IncomingTransfers Return a list of incoming transfers to the wallet.
func (c *Client) IncomingTransfers(req *IncomingTransfersRequest) (*IncomingTransfersResponse, error) {
	resp := &IncomingTransfersResponse{}
	err := c.do("incoming_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
