package walletrpc

type SetTxNotesRequest struct {
	// Txids transaction ids
	Txids []string `json:"txids"`

	// Notes notes for the transactions
	Notes []string `json:"notes"`
}

// SetTxNotes Set arbitrary string notes for transactions.
func (c *Client) SetTxNotes(req *SetTxNotesRequest) error {
	err := c.Do("set_tx_notes", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
