package walletrpc

import "context"

type GetBulkPaymentsRequest struct {
	// Payment IDs used to find the payments (16 characters hex).
	PaymentIds []string `json:"payment_ids,omitempty"`

	// The block height at which to start looking for payments.
	MinBlockHeight uint64 `json:"min_block_height,omitempty"`
}

type GetBulkPaymentsResponse struct {
	// List of payments:
	Payments []Payment `json:"payments"`
}

// Get a list of incoming payments using a given payment id, or a list of payments ids, from a given height. This method is the preferred method over get_payments because it has the same functionality but is more extendable. Either is fine for looking up transactions by a single payment ID.
func (c *Client) GetBulkPayments(ctx context.Context, req *GetBulkPaymentsRequest) (*GetBulkPaymentsResponse, error) {
	resp := &GetBulkPaymentsResponse{}
	err := c.Do(ctx, "get_bulk_payments", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
