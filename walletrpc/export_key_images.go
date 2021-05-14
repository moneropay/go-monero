package walletrpc

type ExportKeyImagesResponse struct {
	// SignedKeyImages array of signed key images:
	SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
}

// ExportKeyImages Export a signed set of key images.
func (c *Client) ExportKeyImages() (*ExportKeyImagesResponse, error) {
	resp := &ExportKeyImagesResponse{}
	err := c.Do("export_key_images", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
