package walletrpc

type TransferSplitRequest struct {
	// Destinations array of destinations to receive XMR:
	Destinations []Destination `json:"destinations"`

	// Amount Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`

	// Address Destination public address.
	Address string `json:"address"`

	// AccountIndex (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint64 `json:"account_index"`

	// SubaddrIndices (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices"`

	// Mixin Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`

	// RingSize Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`

	// UnlockTime Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// GetTxKeys (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys"`

	// Priority Set a priority for the transactions. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority uint64 `json:"priority"`

	// DoNotRelay (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay"`

	// GetTxHex Return the transactions as hex string after sending
	GetTxHex bool `json:"get_tx_hex"`

	// NewAlgorithm True to use the new transaction construction algorithm, defaults to false.
	NewAlgorithm bool `json:"new_algorithm"`

	// GetTxMetadata Return list of transaction metadata needed to relay the transfer later.
	GetTxMetadata bool `json:"get_tx_metadata"`
}

type TransferSplitResponse struct {
	// TxHashList The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`

	// TxKeyList The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`

	// AmountList The amount transferred for every transaction.
	AmountList []int `json:"amount_list"`

	// FeeList The amount of fees paid for every transaction.
	FeeList []int `json:"fee_list"`

	// TxBlobList The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`

	// TxMetadataList List of transaction metadata needed to relay the transactions later.
	TxMetadataList []string `json:"tx_metadata_list"`

	// MultisigTxset The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`

	// UnsignedTxset Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// TransferSplit Same as transfer, but can split into more than one tx if necessary.
func (c *Client) TransferSplit(req *TransferSplitRequest) (*TransferSplitResponse, error) {
	resp := &TransferSplitResponse{}
	err := c.Do("transfer_split", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
