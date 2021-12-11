package walletrpc

import "context"

type GetAddressRequest struct {
	//  Return subaddresses for this account.
	AccountIndex uint64 `json:"account_index"`

	// (Optional) List of subaddresses to return from an account.
	AddressIndex []uint64 `json:"address_index,omitempty"`
}

type GetAddressResponse struct {
	// The 95-character hex address string of the monero-wallet-rpc in session.
	Address string `json:"address"`

	// Addresses
	Addresses []Address `json:"addresses"`
}

// Return the wallet's addresses for an account. Optionally filter for specific set of subaddresses.
func (c *Client) GetAddress(ctx context.Context, req *GetAddressRequest) (*GetAddressResponse, error) {
	resp := &GetAddressResponse{}
	err := c.Do(ctx, "get_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
