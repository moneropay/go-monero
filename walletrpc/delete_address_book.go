package walletrpc

import "context"

type DeleteAddressBookRequest struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// Delete an entry from the address book.
func (c *Client) DeleteAddressBook(ctx context.Context, req *DeleteAddressBookRequest) error {
	return c.Do(ctx, "delete_address_book", &req, nil)
}
