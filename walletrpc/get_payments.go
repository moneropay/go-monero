package walletrpc

import "context"

type GetPaymentsRequest struct {
	// Payment ID used to find the payments (16 characters hex).
	PaymentId string `json:"payment_id"`
}

type GetPaymentsResponse struct {
	// List of payments
	Payments []Payment `json:"payments"`
}

// Get a list of incoming payments using a given payment id.
func (c *Client) GetPayments(ctx context.Context, req *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	resp := &GetPaymentsResponse{}
	err := c.Do(ctx, "get_payments", &req, resp)
	return resp, err
}
