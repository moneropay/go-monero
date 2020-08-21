package walletrpc

type MakeUriRequest struct {
	// Address Wallet address
	Address string `json:"address"`

	// Amount (optional) the integer amount to receive, in atomic units
	Amount uint64 `json:"amount"`

	// PaymentId (optional) 16 or 64 character hexadecimal payment id
	PaymentId string `json:"payment_id"`

	// RecipientName (optional) name of the payment recipient
	RecipientName string `json:"recipient_name"`

	// TxDescription (optional) Description of the reason for the tx
	TxDescription string `json:"tx_description"`
}

type MakeUriResponse struct {
	// Uri This contains all the payment input information as a properly formatted payment URI
	Uri string `json:"uri"`
}

// MakeUri Create a payment URI using the official URI spec.
func (c *Client) MakeUri(req *MakeUriRequest) (*MakeUriResponse, error) {
	resp := &MakeUriResponse{}
	err := c.do("make_uri", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
