package walletrpc

type SignTransferRequest struct {
	// UnsignedTxset Set of unsigned tx returned by "transfer" or "transfer_split" methods.
	UnsignedTxset string `json:"unsigned_txset"`

	// ExportRaw (Optional) If true, return the raw transaction data. (Defaults to false)
	ExportRaw bool `json:"export_raw"`
}

type SignTransferResponse struct {
	// SignedTxset Set of signed tx to be used for submitting transfer.
	SignedTxset string `json:"signed_txset"`

	// TxHashList The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`

	// TxRawList The tx raw data of every transaction.
	TxRawList []string `json:"tx_raw_list"`
}

// SignTransfer Sign a transaction created on a read-only wallet (in cold-signing process)
func (c *Client) SignTransfer(req *SignTransferRequest) (*SignTransferResponse, error) {
	resp := &SignTransferResponse{}
	err := c.Do("sign_transfer", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
