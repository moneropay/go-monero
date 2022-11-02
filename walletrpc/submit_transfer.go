package walletrpc

import "context"

type SubmitTransferRequest struct {
	// Set of signed tx returned by "sign_transfer"
	TxDataHex string `json:"tx_data_hex"`
}

type SubmitTransferResponse struct {
	// The tx hashes of every transaction.
	TxHashList []string `json:"tx_hash_list"`
}

// Submit a previously signed transaction on a read-only wallet (in cold-signing process).
func (c *Client) SubmitTransfer(ctx context.Context, req *SubmitTransferRequest) (*SubmitTransferResponse, error) {
	resp := &SubmitTransferResponse{}
	err := c.Do(ctx, "submit_transfer", &req, resp)
	return resp, err
}
