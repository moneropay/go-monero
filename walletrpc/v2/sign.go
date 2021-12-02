package walletrpc

import "context"

type SignRequest struct {
	// Anything you need to sign.
	Data string `json:"data"`
}

type SignResponse struct {
	// Signature generated against the "data" and the account public address.
	Signature string `json:"signature"`
}

// Sign a string.
func (c *Client) Sign(ctx context.Context, req *SignRequest) (*SignResponse, error) {
	resp := &SignResponse{}
	err := c.Do(ctx, "sign", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
