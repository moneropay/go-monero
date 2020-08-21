package walletrpc

type GetBalanceRequest struct {
	// account_index - unsigned int; Return subaddresses for this account.
	AccountIndex uint `json:"account_index"`
	// address_indices - array of unsigned int; (Optional) List of subaddresses to return from an account.
	AddressIndices []uint `json:"address_indices"`
}

type GetBalanceResponse struct {
	Balance              uint64           `json:"balance"`
	UnlockedBalance      uint64           `json:"unlocked_balance"`
	MultisigImportNeeded bool             `json:"multisig_import_needed"`
	PerSubaddress        []SubaddressInfo `json:"per_subaddress"`
}

type CreateAddressRequest struct {
	AccountIndex uint64 `json:"account_index"`
	Label        string `json:"label"`
}

type CreateAddressResponse struct {
	Address      string `json:"address"`
	AddressIndex uint64 `json:"address_index"`
}

type CreateAccountResponse struct {
	AccountIndex uint64 `json:"account_index"`
	Address      string `json:"address"`
}

type SubaddressInfo struct {
	AddressIndex      uint64 `json:"address_index"`
	Address           string `json:"address"`
	Balance           uint64 `json:"balance"`
	UnlockedBalance   uint64 `json:"unlocked_balance"`
	Label             string `json:"label"`
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`
}

type GetAddressRequest struct {
	// account_index - unsigned int; Return subaddresses for this account.
	AccountIndex uint `json:"account_index"`
	// address_index - array of unsigned int; (Optional) List of subaddresses to return from an account.
	AddressIndex []uint `json:"address_index"`
}

type Addresses struct {
	// address string; The 95-character hex (sub)address string.
	Address string `json:"address"`
	// label string; Label of the (sub)address
	Label string `json:"label"`
	// address_index unsigned int; index of the subaddress
	AddressIndex uint `json:"address_index"`
	// used boolean; states if the (sub)address has already received funds
	Used bool `json:"used"`
}

type GetAddressResponse struct {
	// address - string; The 95-character hex address string of the monero-wallet-rpc in session.
	Address   string      `json:"address"`
	Addresses []Addresses `json:"addresses"`
}

// TransferRequest is the request body of the Transfer client rpc call.
type TransferRequest struct {
	// Destinations - array of destinations to receive XMR:
	Destinations []Destination `json:"destinations"`
	// account_index - unsigned int; (Optional) Transfer from this account index. (Defaults to 0)
	AccountIndex uint `json:"account_index"`
	// subaddr_indices - array of unsigned int; (Optional) Transfer from this set of subaddresses. (Defaults to empty - all indices)
	SubaddrIndices []uint `json:"subaddr_indices"`
	// priority - unsigned int; Set a priority for the transaction.
	// Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
	Priority Priority `json:"priority"`
	// Mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
	Mixin    uint64 `json:"mixin"`
	RingSize uint   `json:"ring_size"`
	// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
	PaymentId string `json:"payment_id,omitempty"`
	// get_tx_key - boolean; (Optional) Return the transaction key after sending.
	GetTxKey bool `json:"get_tx_key"`
	// do_not_relay - boolean; (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// get_tx_hex - boolean; Return the transaction as hex string after sending
	GetTxHex bool `json:"get_tx_hex,omitempty"`
	// get_tx_metadata - boolean; Return the metadata needed to relay the transaction. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata"`
	// Fee - unsigned int; Ignored, will be automatically calculated.
	// Fee uint64 `json:"fee,omitempty"`
}

// Destination to receive XMR
type Destination struct {
	// Amount - unsigned int; Amount to send to each destination, in atomic units.
	Amount uint64 `json:"amount"`
	// Address - string; Destination public address.
	Address string `json:"address"`
}

// TransferResponse is the successful output of a Transfer()
type TransferResponse struct {
	// amount - Amount transferred for the transaction.
	Amount uint64 `json:"amount"`
	// fee - Integer value of the fee charged for the txn.
	Fee uint64 `json:"fee"`
	//multisig_txset - Set of multisig transactions in the process of being signed (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// tx_blob - Transaction as hex string if get_tx_hex is true
	TxBlob string `json:"tx_blob,omitempty"`
	// tx_hash - String for the publically searchable transaction hash
	TxHash string `json:"tx_hash"`
	// tx_key - String for the transaction key if get_tx_key is true, otherwise, blank string.
	TxKey string `json:"tx_key,omitempty"`
	// get_tx_metadata - string; Return the metadata needed to relay the transaction. (Defaults to false)
	TxMetadata string `json:"tx_metadata"`
	// unsigned_txset - String. Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

// TransferSplitResponse is the successful output of a TransferSplit()
type TransferSplitResponse struct {
	// fee_list - array of: integer. The amount of fees paid for every transaction.
	FeeList []uint64 `json:"fee_list"`
	// tx_hash_list - array of: string. The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// tx_blob_list - array of: string. The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// amount_list - array of: integer. The amount transferred for every transaction..
	AmountList []uint64 `json:"amount_list"`
	// tx_key_list - array of: string. The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
}

// SweepAllRequest is the struct to send all unlocked balance to an address.
type SweepAllRequest struct {
	// address - string; Destination public address.
	Address string `json:"address"`
	// priority - unsigned int; (Optional)
	Priority Priority `json:"priority,omitempty"`
	// mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin"`
	// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
	PaymentID string `json:"payment_id,omitempty"`
	// get_tx_keys - boolean; (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// below_amount - unsigned int; (Optional)
	BelowAmount uint64 `json:"below_amount"`
	// do_not_relay - boolean; (Optional)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// get_tx_hex - boolean; (Optional) return the transactions as hex encoded string.
	GetTxHex bool `json:"get_tx_hex,omitempty"`
}

// SweepAllResponse is a typical response of a SweepAllRequest
type SweepAllResponse struct {
	// tx_hash_list - array of: string. The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
	// tx_blob_list - array of: string. The tx as hex string for every transaction.
	TxBlobList []string `json:"tx_blob_list"`
	// tx_key_list - array of: string. The transaction keys for every transaction.
	TxKeyList []string `json:"tx_key_list"`
}

// SweepSingleRequest is the struct to send single unlocked balance to an address.
type SweepSingleRequest struct {
	// address - string; Destination public address.
	Address string `json:"address"`
	// account_index - unsigned int; Sweep transactions from this account.
	AccountIndex uint64 `json:"account_index"`
	// priority - unsigned int; (Optional)
	Priority Priority `json:"priority,omitempty"`
	// mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
	Mixin uint64 `json:"mixin,omitempty"`
	// RingSize - unsigned int; Number of outputs to mix in the transaction (this output + N decoys from the blockchain).
	RingSize uint64 `json:"ring_size,omitempty"`
	// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
	UnlockTime uint64 `json:"unlock_time"`
	// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
	PaymentId string `json:"payment_id,omitempty"`
	// get_tx_keys - boolean; (Optional) Return the transaction keys after sending.
	GetTxKeys bool `json:"get_tx_keys,omitempty"`
	// get_tx_metadata - boolean; (Optional) return the transaction metadata as a string. (Defaults to false)
	GetTxMetadata bool `json:"get_tx_metadata,omitempty"`
	//  string; Key image of specific output to sweep.
	KeyImage string `json:"key_image"`
	// below_amount - unsigned int; (Optional)
	BelowAmount uint64 `json:"below_amount"`
	// do_not_relay - boolean; (Optional)
	DoNotRelay bool `json:"do_not_relay,omitempty"`
	// get_tx_hex - boolean; (Optional) return the transactions as hex encoded string.
	GetTxHex bool `json:"get_tx_hex,omitempty"`
}

// SweepSingleResponse is a typical response of a SweepSingleRequest
type SweepSingleResponse struct {
	// tx_hash_list - string. The tx hashes of every transaction.
	TxHash string `json:"tx_hash"`
	// tx_key_list - string. The transaction keys for every transaction.
	TxKey string `json:"tx_key"`
	// amount - integer. The amount transferred for every transaction.
	Amount uint64 `json:"amount"`
	// fee - integer. The amount of fees paid for every transaction.
	Fee uint64 `json:"fee"`
	// tx_blob - string. The tx as hex string for every transaction.
	TxBlob string `json:"tx_blob"`
	// tx_metadata - string. List of transaction metadata needed to relay the transactions later.
	TxMetadata string `json:"tx_metadata"`
	// multisig_txset - string. The set of signing keys used in a multisig transaction (empty for non-multisig).
	MultisigTxset string `json:"multisig_txset"`
	// unsigned_txset - string. Set of unsigned tx for cold-signing purposes.
	UnsignedTxset string `json:"unsigned_txset"`
}

type SubAddrIndex struct {
	Major uint `json:"major"`
	Minor uint `json:"minor"`
}

// Payment ...
type Payment struct {
	PaymentId    string       `json:"payment_id"`
	TxHash       string       `json:"tx_hash"`
	Amount       uint64       `json:"amount"`
	BlockHeight  uint64       `json:"block_height"`
	UnlockTime   uint64       `json:"unlock_time"`
	SubaddrIndex SubAddrIndex `json:"subaddr_index"`
	Address      string       `json:"address"`
}

// GetTransfersRequest = GetTransfers body
type GetTransfersRequest struct {
	In             bool     `json:"in"`
	Out            bool     `json:"out"`
	Pending        bool     `json:"pending"`
	Failed         bool     `json:"failed"`
	Pool           bool     `json:"pool"`
	FilterByHeight bool     `json:"filter_by_height"`
	MinHeight      uint64   `json:"min_height"`
	MaxHeight      uint64   `json:"max_height"`
	AccountIndex   uint64   `json:"account_index"`
	SubAddrIndex   []uint64 `json:"subaddr_indices"`
}

// GetTransfersResponse = GetTransfers output
type GetTransfersResponse struct {
	In      []Transfer `json:"in"`
	Out     []Transfer `json:"out"`
	Pending []Transfer `json:"pending"`
	Failed  []Transfer `json:"failed"`
	Pool    []Transfer `json:"pool"`
}

type Transfer struct {
	Address                         string       `json:"address"`
	Amount                          uint64       `json:"amount"`
	Confirmations                   uint64       `json:"confirmations"`
	DoubleSpendSeen                 bool         `json:"double_spend_seen"`
	Fee                             uint64       `json:"fee"`
	Height                          uint64       `json:"height"`
	Note                            string       `json:"note"`
	PaymentId                       string       `json:"payment_id"`
	SubaddrIndex                    SubAddrIndex `json:"subaddr_index"`
	SuggestedConfirmationsThreshold uint         `json:"suggested_confirmations_threshold"`
	Timestamp                       uint64       `json:"timestamp"`
	TxId                            string       `json:"txid"`
	Type                            string       `json:"type"`
	UnlockTime                      uint64       `json:"unlock_time"`
}

type GetTransferByTxidRequest struct {
	TxId         string `json:"txid"`
	AccountIndex uint   `json:"account_index,omitempty"`
}

// Transfer is the transfer data of
type GetTransferByTxidResponse struct {
	Transfer Transfer `json:"transfer"`
}

type GetIncomingTransferRequest struct {
	TransferType   GetTransferType `json:"transfer_type"`
	AccountIndex   uint            `json:"account_index"`
	SubaddrIndices SubAddrIndex    `json:"subaddr_indices"`
	Verbose        bool            `json:"verbose"`
}

// IncTransfer is returned by IncomingTransfers
type IncTransfer struct {
	Amount uint64 `json:"amount"`
	// Mostly internal use, can be ignored by most users.
	GlobalIndex  uint64 `json:"global_index"`
	KeyImage     string `json:"key_image"`
	Spent        bool   `json:"spent"`
	SubaddrIndex uint   `json:"subaddr_index"`
	// Several incoming transfers may share the same hash
	// if they were in the same transaction.
	TxHash string `json:"tx_hash"`
	TxSize uint64 `json:"tx_size"`
}

type GetIncomingTransferResponse struct {
	Transfer []IncTransfer `json:"transfer"`
}

// URIDef is the skeleton of the MakeURI()
type URIDef struct {
	// address - wallet address string
	Address string `json:"address"`
	// amount (optional) - the integer amount to receive, in atomic units
	Amount uint64 `json:"amount,omitempty"`
	// payment_id (optional) - 16 or 64 character hexadecimal payment id string
	PaymentID string `json:"payment_id,omitempty"`
	// recipient_name (optional) - string name of the payment recipient
	RecipientName string `json:"recipient_name,omitempty"`
	// tx_description (optional) - string describing the reason for the tx
	TxDescription string `json:"tx_description,omitempty"`
}

// SignedKeyImage - The key image is an alternate public key computed on a second base point,
// Hp(P), instead of G. It is required in traceable ring signature construction to ensure
// multiple signatures with the same real key are linked.
type SignedKeyImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

// ImportKeyImageResponse is the result of ImportKeyImages()
type ImportKeyImageResponse struct {
	Height  uint64 `json:"height"`
	Spent   uint64 `json:"spent"`
	Unspent uint64 `json:"unspent"`
}

// AddressBookEntry Address
type AddressBookEntry struct {
	Address     string `json:"address"`
	Description string `json:"description,omitempty"`
	Index       uint64 `json:"index,omitempty"`
	PaymentId   string `json:"payment_id,omitempty"`
}
