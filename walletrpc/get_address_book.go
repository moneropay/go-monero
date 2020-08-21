package walletrpc

type GetAddressBookRequest struct {
	// Entries indices of the requested address book entries
	Entries []uint64 `json:"entries"`
}

type GetAddressBookResponse struct {
	// Entries array of entries:
	Entries []Entry `json:"entries"`
}

// GetAddressBook Retrieves entries from the address book.
func (c *Client) GetAddressBook(req *GetAddressBookRequest) (*GetAddressBookResponse, error) {
	resp := &GetAddressBookResponse{}
	err := c.do("get_address_book", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
