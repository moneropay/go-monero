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

	// Used states if the (sub)address has already received funds.
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

	// Destination public address.
	Address string `json:"address"`
}

type Payment struct {
	// Payment ID matching the input parameter.
	PaymentId string `json:"payment_id"`

	// Transaction hash used as the transaction ID.
	TxHash string `json:"tx_hash"`

	// Amount for this payment.
	Amount uint64 `json:"amount"`

	// Height of the block that first confirmed this payment.
	BlockHeight uint64 `json:"block_height"`

	// Time (in block height) until this payment is safe to spend.
	UnlockTime uint64 `json:"unlock_time"`

	// SubaddrIndex index:
	SubaddrIndex SubaddressIndex `json:"subaddr_index"`

	// Account index for the subaddress.
	Major uint64 `json:"major"`

	// Index of the subaddress in the account.
	Minor uint64 `json:"minor"`

	// Address receiving the payment; Base58 representation of the public keys.
	Address string `json:"address"`
}

type Transfer struct {
	// Public address of the transfer.
	Address string `json:"address"`

	// Amount of this transfer.
	Amount uint64 `json:"amount"`

	// Number of blocks mined since the block containing this transaction (or block height at which the transaction should be added to a block if not yet confirmed).
	Confirmations uint64 `json:"confirmations"`

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

type SignedKeyImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

type Entry struct {
	// Public address of the entry.
	Address string `json:"address"`

	// Description of this address entry.
	Description string `json:"description"`

	Index uint64 `json:"index"`
}
