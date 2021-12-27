package walletrpc

import "context"

type SweepDustRequest struct {
	// (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`

	// (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`

	// (Optional) Return the transactions as hex string after sending. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`

	// (Optional) Return list of transaction metadata needed to relay the transfer later. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

type SweepDustResponse struct {
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

// Send all dust outputs back to the wallet's, to make them easier to spend (and mix).
func (c *Client) SweepDust(ctx context.Context, req *SweepDustRequest) (*SweepDustResponse, error) {
	resp := &SweepDustResponse{}
	err := c.Do(ctx, "sweep_dust", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
