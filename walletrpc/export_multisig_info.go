package walletrpc

type ExportMultisigInfoResponse struct {
	// Info Multisig info in hex format for other participants.
	Info string `json:"info"`
}

// ExportMultisigInfo Export multisig info for other participants.
func (c *Client) ExportMultisigInfo() (*ExportMultisigInfoResponse, error) {
	resp := &ExportMultisigInfoResponse{}
	err := c.do("export_multisig_info", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
