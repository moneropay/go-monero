package walletrpc

import "context"

type GetTxNotesRequest struct {
	// Transaction ids
	Txids []string `json:"txids"`
}

type GetTxNotesResponse struct {
	// Notes for the transactions
	Notes []string `json:"notes"`
}

// Get string notes for transactions.
func (c *Client) GetTxNotes(ctx context.Context, req *GetTxNotesRequest) (*GetTxNotesResponse, error) {
	resp := &GetTxNotesResponse{}
	err := c.Do(ctx, "get_tx_notes", &req, resp)
	return resp, err
}
