package walletrpc

import "context"

type AddAddressBookRequest struct {
	Address     string `json:"address"`
	PaymentId   string `json:"payment_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type AddAddressBookResponse struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// Add an entry to the address book.
func (c *Client) AddAddressBook(ctx context.Context, req *AddAddressBookRequest) (*AddAddressBookResponse, error) {
	resp := &AddAddressBookResponse{}
	err := c.Do(ctx, "add_address_book", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
