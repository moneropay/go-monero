package walletrpc

import "context"

type TransferSplitRequest struct {
	// Array of destinations to receive XMR.
	Destinations []Destination `json:"destinations"`

	// Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`

	// Destination public address.
	Address string `json:"address"`

	// (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint64 `json:"account_index,omitempty"`

	// (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices,omitempty"`

	// Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`

	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`

	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`

	// Set a priority for the transactions. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority uint64 `json:"priority"`

	// (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`

	// Return the transactions as hex string after sending
	GetTxHex bool `json:"get_tx_hex"`

	// True to use the new transaction construction algorithm, defaults to false.
	NewAlgorithm bool `json:"new_algorithm"`

	// Return list of transaction metadata needed to relay the transfer later.
	GetTxMetadata bool `json:"get_tx_metadata"`
}

type TransferSplitResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`

	// The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`

	// The amount transferred for every transaction.
	AmountList []int `json:"amount_list"`

	// The amount of fees paid for every transaction.
	FeeList []int `json:"fee_list"`

	// The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`

	// List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`

	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`

	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// Same as transfer, but can split into more than one tx if necessary.
func (c *Client) TransferSplit(ctx context.Context, req *TransferSplitRequest) (*TransferSplitResponse, error) {
	resp := &TransferSplitResponse{}
	err := c.Do(ctx, "transfer_split", &req, resp)
	return resp, err
}
