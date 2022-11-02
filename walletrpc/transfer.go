package walletrpc

import "context"

type TransferRequest struct {
	// Array of destinations to receive XMR.
	Destinations []Destination `json:"destinations"`

	// (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`

	// (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`

	// Set a priority for the transaction. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority Priority `json:"priority"`

	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`

	// Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`

	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// (Optional) Return the transaction key after sending.
	GetTxKey bool `json:"get_tx_key,omitempty"`

	// (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`

	// Return the transaction as hex string after sending (Defaults to false)
	GetTxHex bool `json:"get_tx_hex"`

	// Return the metadata needed to relay the transaction. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
}

type TransferResponse struct {
	// Amount transferred for the transaction.
	Amount uint64 `json:"amount"`

	// Value of the fee charged for the txn.
	Fee uint64 `json:"fee"`

	// Set of multisig transactions in the process of being signed (empty for non-multisig).
	MultisigTxset interface{} `json:"multisig_txset"`

	// Raw transaction represented as hex string, if get_tx_hex is true.
	TxBlob string `json:"tx_blob"`

	// The publically searchable transaction hash.
	TxHash string `json:"tx_hash"`

	// The transaction key if get_tx_key is true, otherwise, blank string.
	TxKey string `json:"tx_key"`

	// Set of transaction metadata needed to relay this transfer later, if get_tx_metadata is true.
	TxMetadata string `json:"tx_metadata"`

	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// Send monero to a number of recipients.
func (c *Client) Transfer(ctx context.Context, req *TransferRequest) (*TransferResponse, error) {
	resp := &TransferResponse{}
	err := c.Do(ctx, "transfer", &req, resp)
	return resp, err
}
