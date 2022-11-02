package walletrpc

import "context"

type ImportMultisigInfoRequest struct {
	// List of multisig info in hex format from other participants.
	Info []string `json:"info"`
}

type ImportMultisigInfoResponse struct {
	// Number of outputs signed with those multisig info.
	NOutputs uint64 `json:"n_outputs"`
}

// Import multisig info from other participants.
func (c *Client) ImportMultisigInfo(ctx context.Context, req *ImportMultisigInfoRequest) (*ImportMultisigInfoResponse, error) {
	resp := &ImportMultisigInfoResponse{}
	err := c.Do(ctx, "import_multisig_info", &req, resp)
	return resp, err
}
