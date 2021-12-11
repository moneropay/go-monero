package walletrpc

import "context"

type GetAddressIndexRequest struct {
	// (sub)address to look for.
	Address string `json:"address"`
}

type GetAddressIndexResponse struct {
	// Subaddress information.
	Index SubaddressIndex `json:"index"`
}

// Get account and address indexes from a specific (sub)address.
func (c *Client) GetAddressIndex(ctx context.Context, req *GetAddressIndexRequest) (*GetAddressIndexResponse, error) {
	resp := &GetAddressIndexResponse{}
	err := c.Do(ctx, "get_address_index", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
