package walletrpc

type SweepAllRequest struct {
	// Address Destination public address.
	Address string `json:"address"`

	// AccountIndex Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`

	// SubaddrIndices (Optional) Sweep from this set of subaddresses in the account.
	SubaddrIndices []uint64 `json:"subaddr_indices"`

	// Priority (Optional) Priority for sending the sweep transfer, partially determines fee.
	Priority uint64 `json:"priority"`

	// Mixin Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`

	// RingSize Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size"`

	// UnlockTime Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// GetTxKeys (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys"`

	// BelowAmount (Optional) Include outputs below this amount.
	BelowAmount uint64 `json:"below_amount"`

	// DoNotRelay (Optional) If true, do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay"`

	// GetTxHex (Optional) return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex"`

	// GetTxMetadata (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
}

type SweepAllResponse struct {
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

// SweepAll Send all unlocked balance to an address.
func (c *Client) SweepAll(req *SweepAllRequest) (*SweepAllResponse, error) {
	resp := &SweepAllResponse{}
	err := c.do("sweep_all", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
