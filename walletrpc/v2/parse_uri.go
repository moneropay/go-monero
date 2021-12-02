package walletrpc

import "context"

type ParseUriRequest struct {
	// This contains all the payment input information as a properly formatted payment URI.
	Uri string `json:"uri"`
}

type ParseUriResponse struct {
	// Uri containing payment information.
	Uri UriPaymentInfo `json:"uri"`
}

type UriPaymentInfo struct {
	// Wallet address.
	Address string `json:"address"`

	// Amount to receive, in atomic units (0 if not provided)
	Amount uint64 `json:"amount"`

	// 16 or 64 character hexadecimal payment id (empty if not provided).
	PaymentId string `json:"payment_id"`

	// Name of the payment recipient (empty if not provided).
	RecipientName string `json:"recipient_name"`

	// Description of the reason for the tx (empty if not provided).
	TxDescription string `json:"tx_description"`
}

// Parse a payment URI to get payment information.
func (c *Client) ParseUri(ctx context.Context, req *ParseUriRequest) (*ParseUriResponse, error) {
	resp := &ParseUriResponse{}
	err := c.Do(ctx, "parse_uri", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
