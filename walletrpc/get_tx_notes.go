package walletrpc

type GetTxNotesRequest struct {
	// Txids transaction ids
	Txids []string `json:"txids"`
}

type GetTxNotesResponse struct {
	// Notes notes for the transactions
	Notes []string `json:"notes"`
}

// GetTxNotes Get string notes for transactions.
func (c *Client) GetTxNotes(req *GetTxNotesRequest) (*GetTxNotesResponse, error) {
	resp := &GetTxNotesResponse{}
	err := c.do("get_tx_notes", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
