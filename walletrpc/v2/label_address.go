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
	err := c.Do(ctx, "label_address", &req, nil)
	if err != nil {
		return err
	}
	return nil
}
