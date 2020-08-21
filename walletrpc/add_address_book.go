package walletrpc

type AddAddressBookRequest struct {
	Address     string `json:"address"`
	PaymentId   string `json:"payment_id"`
	Description string `json:"description"`
}

type AddAddressBookResponse struct {
	// Index The index of the address book entry.
	Index uint64 `json:"index"`
}

// AddAddressBook Add an entry to the address book.
func (c *Client) AddAddressBook(req *AddAddressBookRequest) (*AddAddressBookResponse, error) {
	resp := &AddAddressBookResponse{}
	err := c.do("add_address_book", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
