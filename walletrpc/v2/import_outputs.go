package walletrpc

import "context"

type ImportOutputsRequest struct {
	// Wallet outputs in hex format.
	OutputsDataHex int64 `json:"outputs_data_hex"`
}

type ImportOutputsResponse struct {
	// Number of outputs imported.
	NumImported uint64 `json:"num_imported"`
}

// Import outputs in hex format.
func (c *Client) ImportOutputs(ctx context.Context, req *ImportOutputsRequest) (*ImportOutputsResponse, error) {
	resp := &ImportOutputsResponse{}
	err := c.Do(ctx, "import_outputs", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
