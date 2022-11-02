package walletrpc

import "context"

type EditAddressBookRequest struct {
	// Index of the address book entry to edit.
	Index uint64 `json:"index"`

	// If true, set the address for this entry to the value of "address".
	SetAddress bool `json:"set_address"`

	// (Optional) The 95-character public address to set.
	Address string `json:"address,omitempty"`

	// If true, set the description for this entry to the value of "description".
	SetDescription bool `json:"set_description"`

	// (Optional) Human-readable description for this entry.
	Description string `json:"description,omitempty"`

	// If true, set the payment ID for this entry to the value of "payment_id".
	SetPaymentId bool `json:"set_payment_id"`

	// (Optional) Payment ID for this address.
	PaymentId string `json:"payment_id,omitempty"`
}

// Edit an existing address book entry.
func (c *Client) EditAddressBook(ctx context.Context, req *EditAddressBookRequest) error {
	return c.Do(ctx, "edit_address_book", &req, nil)
}
