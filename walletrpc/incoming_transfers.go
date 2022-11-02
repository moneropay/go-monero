package walletrpc

import "context"

type IncomingTransfersRequest struct {
	// "all": all the transfers, "available": only transfers which are not yet spent, OR "unavailable": only transfers which are already spent.
	TransferType string `json:"transfer_type"`

	// (Optional) Return transfers for this account. (defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`

	// (Optional) Return transfers sent to these subaddresses.
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`
}

type IncomingTransfer struct {
	// Amount of this transfer.
	Amount uint64 `json:"amount"`

	// Mostly internal use, can be ignored by most users.
	GlobalIndex uint64 `json:"global_index"`

	// Key image for the incoming transfer's unspent output (empty unless verbose is true).
	KeyImage string `json:"key_image"`

	// Indicates if this transfer has been spent.
	Spent bool `json:"spent"`

	// Subaddress index for incoming transfer.
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`

	// Several incoming transfers may share the same hash if they were in the same transaction.
	TxHash string `json:"tx_hash"`

	// Size of transaction in bytes.
	TxSize uint64 `json:"tx_size"`
}

type IncomingTransfersResponse struct {
	Transfers []IncomingTransfer `json:"transfers"`
}

// Return a list of incoming transfers to the wallet.
func (c *Client) IncomingTransfers(ctx context.Context, req *IncomingTransfersRequest) (*IncomingTransfersResponse, error) {
	resp := &IncomingTransfersResponse{}
	err := c.Do(ctx, "incoming_transfers", &req, resp)
	return resp, err
}
