package walletrpc

import "context"

type SignTransferRequest struct {
	// Set of unsigned tx returned by "transfer" or "transfer_split" methods.
	UnsignedTxset string `json:"unsigned_txset"`

	// (Optional) If true, return the raw transaction data. (Defaults to false)
	ExportRaw bool `json:"export_raw,omitempty"`
}

type SignTransferResponse struct {
	// Set of signed tx to be used for submitting transfer.
	SignedTxset string `json:"signed_txset"`

	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`

	// The tx raw data of every transaction.
	TxRawList []string `json:"tx_raw_list"`
}

// Sign a transaction created on a read-only wallet (in cold-signing process)
func (c *Client) SignTransfer(ctx context.Context, req *SignTransferRequest) (*SignTransferResponse, error) {
	resp := &SignTransferResponse{}
	err := c.Do(ctx, "sign_transfer", &req, resp)
	return resp, err
}
