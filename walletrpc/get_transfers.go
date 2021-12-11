package walletrpc

import "context"

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

type Transfer struct {
	// Public address of the transfer.
	Address string `json:"address"`

	// Amount of this transfer.
	Amount uint64 `json:"amount"`

	// Number of blocks mined since the block containing this transaction (or block height at which the transaction should be added to a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`

	// True if the key image(s) for the transfer have been seen before.
	DoubleSpendSeen bool `json:"double_spend_seen"`

	// Transaction fee for this transfer.
	Fee uint64 `json:"fee"`

	// Height of the first block that confirmed this transfer (0 if not mined yet).
	Height uint64 `json:"height"`

	// Note about this transfer.
	Note string `json:"note"`

	// Payment ID for this transfer.
	PaymentId string `json:"payment_id"`

	// Account and subaddress index.
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`

	// Estimation of the confirmations needed for the transaction to be included in a block.
	SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`

	// POSIX timestamp for when this transfer was first confirmed in a block (or timestamp submission if not mined yet).
	Timestamp uint64 `json:"timestamp"`

	// Transaction ID for this transfer.
	Txid string `json:"txid"`

	// Transfer type.
	Type string `json:"type"`

	// Number of blocks until transfer is safely spendable.
	UnlockTime uint64 `json:"unlock_time"`
}

type GetTransfersResponse struct {
	In      []Transfer `json:"in"`
	Out     []Transfer `json:"out"`
	Pending []Transfer `json:"pending"`
	Failed  []Transfer `json:"failed"`
	Pool    []Transfer `json:"pool"`
}

// Returns a list of transfers.
func (c *Client) GetTransfers(ctx context.Context, req *GetTransfersRequest) (*GetTransfersResponse, error) {
	resp := &GetTransfersResponse{}
	err := c.Do(ctx, "get_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
