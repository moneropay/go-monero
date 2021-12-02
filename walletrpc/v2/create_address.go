package walletrpc

import "context"

type CreateAddressRequest struct {
	// Create a new address for this account.
	AccountIndex uint64 `json:"account_index"`

	// (Optional) Label for the new address.
	Label string `json:"label,omitempty"`
}

type CreateAddressResponse struct {
	// Newly created address. Base58 representation of the public keys.
	Address string `json:"address"`

	// Index of the new address under the input account.
	AddressIndex uint64 `json:"address_index"`
}

// Create a new address for an account. Optionally, label the new address.
func (c *Client) CreateAddress(ctx context.Context, req *CreateAddressRequest) (*CreateAddressResponse, error) {
	resp := &CreateAddressResponse{}
	err := c.Do(ctx, "create_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
