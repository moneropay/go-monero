package walletrpc

type GetTransferByTxidRequest struct {
	// Transaction ID used to find the transfer.
	Txid string `json:"txid"`

	// (Optional) Index of the account to query for the transfer.
	AccountIndex uint64 `json:"account_index,omitempty"`
}

type TransferByTxid struct {
	// Public address of the transfer.
	Address string `json:"address"`

	// Amount of this transfer.
	Amount uint64 `json:"amount"`

	// Number of blocks mined since the block containing this transaction (or block height at which the transaction should be added to a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`

	// Array of JSON objects containing transfer destinations.
	Destinations []Destination `json:"destinations"`

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

type GetTransferByTxidResponse struct {
	// Transfer containing payment information:
	Transfer TransferByTxid `json:"transfer"`
}

// Show information about a transfer to/from this address.
func (c *Client) GetTransferByTxid(req *GetTransferByTxidRequest) (*GetTransferByTxidResponse, error) {
	resp := &GetTransferByTxidResponse{}
	err := c.Do("get_transfer_by_txid", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
