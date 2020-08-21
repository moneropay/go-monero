package walletrpc

type GetLanguagesResponse struct {
	// Languages List of available languages
	Languages []string `json:"languages"`
}

// GetLanguages Get a list of available languages for your wallet's seed.
func (c *Client) GetLanguages() (*GetLanguagesResponse, error) {
	resp := &GetLanguagesResponse{}
	err := c.do("get_languages", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
