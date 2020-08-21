package walletrpc

type ImportKeyImagesRequest struct {
	// SignedKeyImages array of signed key images:
	SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
}

type ImportKeyImagesResponse struct {
	// Height int;
	Height uint64 `json:"height"`

	// Spent Amount (in atomic units) spent from those key images.
	Spent uint64 `json:"spent"`

	// Unspent Amount (in atomic units) still available from those key images.
	Unspent uint64 `json:"unspent"`
}

// ImportKeyImages Import signed key images list and verify their spent status.
func (c *Client) ImportKeyImages(req *ImportKeyImagesRequest) (*ImportKeyImagesResponse, error) {
	resp := &ImportKeyImagesResponse{}
	err := c.do("import_key_images", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
