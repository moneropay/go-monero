package walletrpc

type SetTxNotesRequest struct {
	// Transaction ids
	Txids []string `json:"txids"`

	// Notes for the transactions
	Notes []string `json:"notes"`
}

// Set arbitrary string notes for transactions.
func (c *Client) SetTxNotes(req *SetTxNotesRequest) error {
	err := c.Do("set_tx_notes", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
