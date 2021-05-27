package walletrpc

type ExportOutputsRequest struct {
	// (Optional) If true, export all outputs. Otherwise, export outputs since the last export. (Defaults to false)
	All bool `json:"all,omitemtpy"`
}

type ExportOutputsResponse struct {
	// Wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// Export all outputs in hex format.
func (c *Client) ExportOutputs(req *ExportOutputsRequest) (*ExportOutputsResponse, error) {
	resp := &ExportOutputsResponse{}
	err := c.Do("export_outputs", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
