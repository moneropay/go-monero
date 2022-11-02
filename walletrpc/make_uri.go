package walletrpc

import "context"

type MakeUriRequest struct {
	// Wallet address.
	Address string `json:"address"`

	// (Optional) The integer amount to receive, in atomic units.
	Amount uint64 `json:"amount,omitempty"`

	// (Optional) 16 or 64 character hexadecimal payment id.
	PaymentId string `json:"payment_id,omitempty"`

	// (Optional) Name of the payment recipient.
	RecipientName string `json:"recipient_name,omitempty"`

	// (Optional) Description of the reason for the tx.
	TxDescription string `json:"tx_description,omitempty"`
}

type MakeUriResponse struct {
	// This contains all the payment input information as a properly formatted payment URI.
	Uri string `json:"uri"`
}

// Create a payment URI using the official URI spec.
func (c *Client) MakeUri(ctx context.Context, req *MakeUriRequest) (*MakeUriResponse, error) {
	resp := &MakeUriResponse{}
	err := c.Do(ctx, "make_uri", &req, resp)
	return resp, err
}
