package walletrpc

type CreateAddressRequest struct {
	// AccountIndex  Create a new address for this account.
	AccountIndex uint64 `json:"account_index"`

	// Label  (Optional) Label for the new address.
	Label string `json:"label"`
}

type CreateAddressResponse struct {
	// Address  Newly created address. Base58 representation of the public keys.
	Address string `json:"address"`

	// AddressIndex  Index of the new address under the input account.
	AddressIndex uint64 `json:"address_index"`
}

// CreateAddress Create a new address for an account. Optionally, label the new address.
func (c *Client) CreateAddress(req *CreateAddressRequest) (*CreateAddressResponse, error) {
	resp := &CreateAddressResponse{}
	err := c.Do("create_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
