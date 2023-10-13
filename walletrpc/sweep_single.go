package walletrpc

import "context"

type SweepSingleRequest struct {
	// Destination public address.
	Address string `json:"address"`

	// (Optional) Priority for sending the sweep transfer, partially determines fee.
	Priority Priority `json:"priority,omitempty"`

	// Specify the number of separate outputs of smaller denomination that will be created by sweep operation.
	Outputs uint64 `json:"outputs,omitempty"`

	// Sets ringsize to n (mixin + 1).
	RingSize uint64 `json:"ring_size,omitempty"`

	// Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`

	// (Optional, defaults to a random ID) 16 characters hex encoded.
	PaymentId string `json:"payment_id,omitempty"`

	// (Optional) Return the transaction keys after sending.
	GetTxKey bool `json:"get_tx_key,omitempty"`

	// Key image of specific output to sweep.
	KeyImage string `json:"key_image"`

	// (Optional) If true, Do not relay this sweep transfer. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`

	// (Optional) return the transactions as hex encoded string. (Defaults to false)
	GetTxHex bool `json:"get_tx_hex,omitempty"`

	// (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
}

type SweepSingleResponse struct {
	// The tx hashes of every transaction.
	TxHash []string `json:"tx_hash"`

	// The transaction keys for every transaction.
	TxKey []string `json:"tx_key"`

	// The amount transferred for every transaction.
	Amount []int `json:"amount"`

	// The amount of fees paid for every transaction.
	Fee []int `json:"fee"`

	// Metric used to calculate transaction fee.
	Weight uint64 `json:"weight"`

	// The tx as hex string for every transaction.
	TxBlob []string `json:"tx_blob"`

	// List of transaction metadata needed to relay the transactions later.
	TxMetadata []string `json:"tx_metadata"`

	// The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`

	// Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`

	// Key images of spent outputs.
	SpentKeyImages []KeyImages `json:"spent_key_images"`
}

// Send all of a specific unlocked output to an address.
func (c *Client) SweepSingle(ctx context.Context, req *SweepSingleRequest) (*SweepSingleResponse, error) {
	resp := &SweepSingleResponse{}
	err := c.Do(ctx, "sweep_single", &req, resp)
	return resp, err
}
