package walletrpc

type GetLanguagesResponse struct {
	// List of available languages
	Languages []string `json:"languages"`
}

// Get a list of available languages for your wallet's seed.
func (c *Client) GetLanguages() (*GetLanguagesResponse, error) {
	resp := &GetLanguagesResponse{}
	err := c.Do("get_languages", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
