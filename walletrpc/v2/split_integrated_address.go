package walletrpc

import "context"

type SplitIntegratedAddressRequest struct {
	IntegratedAddress string `json:"integrated_address"`
}

type SplitIntegratedAddressResponse struct {
	// States if the address is a subaddress
	IsSubaddress bool `json:"is_subaddress"`

	// Hex encoded payment.
	Payment string `json:"payment"`

	StandardAddress string `json:"standard_address"`
}

// Retrieve the standard address and payment id corresponding to an integrated address.
func (c *Client) SplitIntegratedAddress(ctx context.Context, req *SplitIntegratedAddressRequest) (*SplitIntegratedAddressResponse, error) {
	resp := &SplitIntegratedAddressResponse{}
	err := c.Do(ctx, "split_integrated_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
