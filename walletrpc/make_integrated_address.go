package walletrpc

type MakeIntegratedAddressRequest struct {
	// StandardAddress (Optional, defaults to primary address) Destination public address.
	StandardAddress string `json:"standard_address"`

	// PaymentId (Optional, defaults to a random ID) 16 characters hex encoded.
	PaymentId string `json:"payment_id"`
}

type MakeIntegratedAddressResponse struct {
	// PaymentId hex encoded;
	PaymentId string `json:"payment_id"`
}

// MakeIntegratedAddress Make an integrated address from the wallet address and a payment id.
func (c *Client) MakeIntegratedAddress(req *MakeIntegratedAddressRequest) (*MakeIntegratedAddressResponse, error) {
	resp := &MakeIntegratedAddressResponse{}
	err := c.Do("make_integrated_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
