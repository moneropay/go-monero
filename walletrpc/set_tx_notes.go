package walletrpc

import "context"

type SetTxNotesRequest struct {
	// Transaction ids
	Txids []string `json:"txids"`

	// Notes for the transactions
	Notes []string `json:"notes"`
}

// Set arbitrary string notes for transactions.
func (c *Client) SetTxNotes(ctx context.Context, req *SetTxNotesRequest) error {
	return c.Do(ctx, "set_tx_notes", &req, nil)
}
