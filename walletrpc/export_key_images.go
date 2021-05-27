package walletrpc

type ExportKeyImagesRequest struct {
	// (Optional) If true, export all key images. Otherwise, export key images since the last export. (default = false)
	All bool `json:"all,omitempty"`
}
type ExportKeyImagesResponse struct {
	// Array of signed key images:
	SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
}

// Export a signed set of key images.
func (c *Client) ExportKeyImages(req *ExportKeyImagesRequest) (*ExportKeyImagesResponse, error) {
	resp := &ExportKeyImagesResponse{}
	err := c.Do("export_key_images", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
