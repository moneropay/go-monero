package walletrpc

type SplitIntegratedAddressRequest struct {
	IntegratedAddress string `json:"integrated_address"`
}

type SplitIntegratedAddressResponse struct {
	// IsSubaddress States if the address is a subaddress
	IsSubaddress bool `json:"is_subaddress"`

	// Payment hex encoded
	Payment string `json:"payment"`

	StandardAddress string `json:"standard_address"`
}

// SplitIntegratedAddress Retrieve the standard address and payment id corresponding to an integrated address.
func (c *Client) SplitIntegratedAddress(req *SplitIntegratedAddressRequest) (*SplitIntegratedAddressResponse, error) {
	resp := &SplitIntegratedAddressResponse{}
	err := c.Do("split_integrated_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
