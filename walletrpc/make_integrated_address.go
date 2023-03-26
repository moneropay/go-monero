package walletrpc

import "context"

type MakeIntegratedAddressRequest struct {
	// (Optional) Destination public address. (Defaults to primary address)
	StandardAddress string `json:"standard_address,omitempty"`

	// (Optional) 16 characters hex encoded. (Defaults to a random ID)
	PaymentId string `json:"payment_id,omitempty"`
}

type MakeIntegratedAddressResponse struct {
	// Hex encoded payment id.
	PaymentId string `json:"payment_id"`

	// Integrated address.
	IntegratedAddress string `json:"integrated_address"`
}

// Make an integrated address from the wallet address and a payment id.
func (c *Client) MakeIntegratedAddress(ctx context.Context, req *MakeIntegratedAddressRequest) (*MakeIntegratedAddressResponse, error) {
	resp := &MakeIntegratedAddressResponse{}
	err := c.Do(ctx, "make_integrated_address", &req, resp)
	return resp, err
}
