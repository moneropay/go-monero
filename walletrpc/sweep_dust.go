package walletrpc

type SweepDustRequest struct {
	// GetTxKeys (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys"`

	// DoNotRelay (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay"`

	// GetTxHex (Optional) Return the transactions as hex string after sending. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex"`

	// GetTxMetadata (Optional) Return list of transaction metadata needed to relay the transfer later. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
}

type SweepDustResponse struct {
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

// SweepDust Send all dust outputs back to the wallet's, to make them easier to spend (and mix).
func (c *Client) SweepDust(req *SweepDustRequest) (*SweepDustResponse, error) {
	resp := &SweepDustResponse{}
	err := c.Do("sweep_dust", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
