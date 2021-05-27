package walletrpc

type GetTxNotesRequest struct {
	// Transaction ids
	Txids []string `json:"txids"`
}

type GetTxNotesResponse struct {
	// Notes for the transactions
	Notes []string `json:"notes"`
}

// Get string notes for transactions.
func (c *Client) GetTxNotes(req *GetTxNotesRequest) (*GetTxNotesResponse, error) {
	resp := &GetTxNotesResponse{}
	err := c.Do("get_tx_notes", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
