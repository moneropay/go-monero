package walletrpc

type GetAddressIndexRequest struct {
	// Address (sub)address to look for.
	Address string `json:"address"`
}

type GetAddressIndexResponse struct {
	// Index subaddress informations
	Index SubaddressIndex `json:"index"`
}

// GetAddressIndex Get account and address indexes from a specific (sub)address
func (c *Client) GetAddressIndex(req *GetAddressIndexRequest) (*GetAddressIndexResponse, error) {
	resp := &GetAddressIndexResponse{}
	err := c.Do("get_address_index", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
