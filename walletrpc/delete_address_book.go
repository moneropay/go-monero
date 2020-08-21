package walletrpc

type DeleteAddressBookRequest struct {
	// Index The index of the address book entry.
	Index uint64 `json:"index"`
}

// DeleteAddressBook Delete an entry from the address book.
func (c *Client) DeleteAddressBook(req *DeleteAddressBookRequest) error {
	err := c.Do("delete_address_book", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
