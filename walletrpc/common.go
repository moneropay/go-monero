package walletrpc

type Address struct {
	// AddressIndex Index of the subaddress in the account.
	AddressIndex uint64 `json:"address_index"`

	// Address at this index. Base58 representation of the public keys.
	Address string `json:"address"`

	// Balance for the subaddress (locked or unlocked).
	Balance uint64 `json:"balance"`

	// UnlockedBalance Unlocked balance for the subaddress.
	UnlockedBalance uint64 `json:"unlocked_balance"`

	// Label of the subaddress.
	Label string `json:"label"`

	// Tag for filtering accounts.
	Tag string `json:"tag"`

	// NumUnspentOutputs Number of unspent outputs available for the subaddress.
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`

	// Used states if the (sub)address has already received funds
	Used bool `json:"used"`
}

type SubaddressIndex struct {
	// Major  Account index.
	Major uint64 `json:"major"`

	// Minor  Address index.
	Minor uint64 `json:"minor"`
}

type Destination struct {
	// Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`

	// Address Destination public address.
	Address string `json:"address"`
}

type Payment struct {
	// PaymentId Payment ID matching the input parameter.
	PaymentId string `json:"payment_id"`

	// TxHash Transaction hash used as the transaction ID.
	TxHash string `json:"tx_hash"`

	// Amount for this payment.
	Amount uint64 `json:"amount"`

	// BlockHeight Height of the block that first confirmed this payment.
	BlockHeight uint64 `json:"block_height"`

	// UnlockTime Time (in block height) until this payment is safe to spend.
	UnlockTime uint64 `json:"unlock_time"`

	// SubaddrIndex index:
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`

	// Major Account index for the subaddress.
	Major uint64 `json:"major"`

	// Minor Index of the subaddress in the account.
	Minor uint64 `json:"minor"`

	// Address Address receiving the payment; Base58 representation of the public keys.
	Address string `json:"address"`
}

type Transfer struct {
	// Amount of this transfer.
	Amount uint64 `json:"amount"`

	// GlobalIndex Mostly internal use, can be ignored by most users.
	GlobalIndex uint64 `json:"global_index"`

	// KeyImage Key image for the incoming transfer's unspent output (empty unless verbose is true).
	KeyImage string `json:"key_image"`

	// Spent Indicates if this transfer has been spent.
	Spent bool `json:"spent"`

	// TxHash Several incoming transfers may share the same hash if they were in the same transaction.
	TxHash string `json:"tx_hash"`

	// TxSize Size of transaction in bytes.
	TxSize uint64 `json:"tx_size"`

	// Address Public address of the transfer.
	Address string `json:"address"`

	// Confirmations Number of block mined since the block containing this transaction (or block height at which the transaction should be added to a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`

	// DoubleSpendSeen True if the key image(s) for the transfer have been seen before.
	DoubleSpendSeen bool `json:"double_spend_seen"`

	// Fee Transaction fee for this transfer.
	Fee uint64 `json:"fee"`

	// Height Height of the first block that confirmed this transfer (0 if not mined yet).
	Height uint64 `json:"height"`

	// Note Note about this transfer.
	Note string `json:"note"`

	// PaymentId Payment ID for this transfer.
	PaymentId string `json:"payment_id"`

	// SubaddrIndex & minor subaddress index:
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`

	// Major Account index for the subaddress.
	Major uint64 `json:"major"`

	// Minor Index of the subaddress under the account.
	Minor uint64 `json:"minor"`

	// SuggestedConfirmationsThreshold Estimation of the confirmations needed for the transaction to be included in a block.
	SuggestedConfirmationsThreshold uint64 `json:"suggested_confirmations_threshold"`

	// Timestamp POSIX timestamp for when this transfer was first confirmed in a block (or timestamp submission if not mined yet).
	Timestamp uint64 `json:"timestamp"`

	// Txid Transaction ID for this transfer.
	Txid string `json:"txid"`

	// Type Transfer type: "in"
	Type string `json:"type"`

	// UnlockTime Number of blocks until transfer is safely spendable.
	UnlockTime uint64 `json:"unlock_time"`
}

type SignedKeyImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

type Entry struct {
	// Address Public address of the entry
	Address string `json:"address"`

	// Description Description of this address entry
	Description string `json:"description"`

	// Index;
	Index uint64 `json:"index"`
}
