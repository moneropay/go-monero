package walletrpc

import "context"

type GetAddressBookRequest struct {
	// Indices of the requested address book entries.
	Entries []uint64 `json:"entries"`
}

type GetAddressBookResponse struct {
	// Array of entries.
	Entries []Entry `json:"entries"`
}

// Retrieves entries from the address book.
func (c *Client) GetAddressBook(ctx context.Context, req *GetAddressBookRequest) (*GetAddressBookResponse, error) {
	resp := &GetAddressBookResponse{}
	err := c.Do(ctx, "get_address_book", &req, resp)
	return resp, err
}
