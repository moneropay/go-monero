package walletrpc

import "context"

type GetLanguagesResponse struct {
	// List of available languages
	Languages []string `json:"languages"`
}

// Get a list of available languages for your wallet's seed.
func (c *Client) GetLanguages(ctx context.Context) (*GetLanguagesResponse, error) {
	resp := &GetLanguagesResponse{}
	err := c.Do(ctx, "get_languages", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
