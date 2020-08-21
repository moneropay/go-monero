package walletrpc

type GetAddressRequest struct {
	// AccountIndex Return subaddresses for this account.
	AccountIndex uint64 `json:"account_index"`

	// AddressIndex (Optional) List of subaddresses to return from an account.
	AddressIndex []uint64 `json:"address_index"`
}

type GetAddressResponse struct {
	// Address The 95-character hex address string of the monero-wallet-rpc in session.
	Address string `json:"address"`

	// Addresses
	Addresses []Address `json:"addresses"`
}

// GetAddress Return the wallet's addresses for an account. Optionally filter for specific set of subaddresses.
func (c *Client) GetAddress(req *GetAddressRequest) (*GetAddressResponse, error) {
	resp := &GetAddressResponse{}
	err := c.do("get_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
