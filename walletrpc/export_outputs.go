package walletrpc

type ExportOutputsResponse struct {
	// OutputsDataHex wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

// ExportOutputs Export all outputs in hex format.
func (c *Client) ExportOutputs() (*ExportOutputsResponse, error) {
	resp := &ExportOutputsResponse{}
	err := c.do("export_outputs", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
