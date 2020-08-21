package walletrpc

type ImportMultisigInfoRequest struct {
	// Info List of multisig info in hex format from other participants.
	Info []string `json:"info"`
}

type ImportMultisigInfoResponse struct {
	// NOutputs Number of outputs signed with those multisig info.
	NOutputs uint64 `json:"n_outputs"`
}

// ImportMultisigInfo Import multisig info from other participants.
func (c *Client) ImportMultisigInfo(req *ImportMultisigInfoRequest) (*ImportMultisigInfoResponse, error) {
	resp := &ImportMultisigInfoResponse{}
	err := c.do("import_multisig_info", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
