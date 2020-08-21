package walletrpc

type TransferRequest struct {
	// Destinations XMR:
	Destinations []Destination `json:"destinations"`

	// AccountIndex (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint64 `json:"account_index"`

	// SubaddrIndices (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint64 `json:"subaddr_indices"`

	// Priority Set a priority for the transaction. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority Priority `json:"priority"`

	// Mixin Number of outputs from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`

	// RingSize Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`

	// UnlockTime Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// GetTxKey (Optional) Return the transaction key after sending.
	GetTxKey bool `json:"get_tx_key"`

	// DoNotRelay (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay"`

	// GetTxHex Return the transaction as hex string after sending (Defaults to false)
	GetTxHex bool `json:"get_tx_hex"`

	// GetTxMetadata Return the metadata needed to relay the transaction. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
}

type TransferResponse struct {
	// Amount transferred for the transaction.
	Amount int64 `json:"amount"`

	// Fee value of the fee charged for the txn.
	Fee int64 `json:"fee"`

	// MultisigTxset (empty for non-multisig).
	MultisigTxset interface{} `json:"multisig_txset"`

	// TxBlob string, if get_tx_hex is true.
	TxBlob string `json:"tx_blob"`

	// TxHash for the publically searchable transaction hash.
	TxHash string `json:"tx_hash"`

	// TxKey for the transaction key if get_tx_key is true, otherwise, blank string.
	TxKey string `json:"tx_key"`

	// TxMetadata Set of transaction metadata needed to relay this transfer later, if get_tx_metadata is true.
	TxMetadata string `json:"tx_metadata"`

	// UnsignedTxset Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// Transfer Send monero to a number of recipients.
func (c *Client) Transfer(req *TransferRequest) (*TransferResponse, error) {
	resp := &TransferResponse{}
	err := c.do("transfer", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
