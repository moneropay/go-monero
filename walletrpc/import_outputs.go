package walletrpc

type ImportOutputsRequest struct {
	// OutputsDataHex wallet outputs in hex format.
	OutputsDataHex string `json:"outputs_data_hex"`
}

type ImportOutputsResponse struct {
	// NumImported number of outputs imported.
	NumImported uint64 `json:"num_imported"`
}

// ImportOutputs Import outputs in hex format.
func (c *Client) ImportOutputs(req *ImportOutputsRequest) (*ImportOutputsResponse, error) {
	resp := &ImportOutputsResponse{}
	err := c.Do("import_outputs", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
