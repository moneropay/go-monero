package walletrpc

type SubmitTransferRequest struct {
	// TxDataHex Set of signed tx returned by "sign_transfer"
	TxDataHex string `json:"tx_data_hex"`
}

type SubmitTransferResponse struct {
	// TxHashList The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
}

// SubmitTransfer Submit a previously signed transaction on a read-only wallet (in cold-signing process).
func (c *Client) SubmitTransfer(req *SubmitTransferRequest) (*SubmitTransferResponse, error) {
	resp := &SubmitTransferResponse{}
	err := c.do("submit_transfer", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
