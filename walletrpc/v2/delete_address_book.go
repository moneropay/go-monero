package walletrpc

import "context"

type DeleteAddressBookRequest struct {
	// The index of the address book entry.
	Index uint64 `json:"index"`
}

// Delete an entry from the address book.
func (c *Client) DeleteAddressBook(ctx context.Context, req *DeleteAddressBookRequest) error {
	err := c.Do(ctx, "delete_address_book", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
