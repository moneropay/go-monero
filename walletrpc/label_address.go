package walletrpc

import "context"

type LabelAddressRequest struct {
	// Struct containing the major & minor address index.
	Index SubaddressIndex `json:"index"`

	// Label for the address.
	Label string `json:"label"`
}

// Label an address.
func (c *Client) LabelAddress(ctx context.Context, req *LabelAddressRequest) error {
	return c.Do(ctx, "label_address", &req, nil)
}
