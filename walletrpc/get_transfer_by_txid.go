package walletrpc

type GetTransferByTxidRequest struct {
	// Txid Transaction ID used to find the transfer.
	Txid string `json:"txid"`

	// AccountIndex (Optional) Index of the account to query for the transfer.
	AccountIndex uint64 `json:"account_index"`
}

type GetTransferByTxidResponse struct {
	// Transfer containing payment information:
	Transfer Transfer `json:"transfer"`
}

// GetTransferByTxid Show information about a transfer to/from this address.
func (c *Client) GetTransferByTxid(req *GetTransferByTxidRequest) (*GetTransferByTxidResponse, error) {
	resp := &GetTransferByTxidResponse{}
	err := c.Do("get_transfer_by_txid", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
