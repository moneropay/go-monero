package walletrpc

type ValidateAddressRequest struct {
	// Address The address to validate.
	Address string `json:"address"`

	// AnyNetType (Optional); If true, consider addresses belonging to any of the three Monero networks (mainnet, stagenet, and testnet) valid. Otherwise, only consider an address valid if it belongs to the network on which the rpc-wallet's current daemon is running (Defaults to false).
	AnyNetType bool `json:"any_net_type"`

	// AllowOpenalias (Optional); If true, consider OpenAlias-formatted addresses valid (Defaults to false).
	AllowOpenalias bool `json:"allow_openalias"`
}

type ValidateAddressResponse struct {
	// Valid True if the input address is a valid Monero address.
	Valid bool `json:"valid"`

	// Integrated True if the given address is an integrated address.
	Integrated bool `json:"integrated"`

	// Subaddress True if the given address is a subaddress
	Subaddress bool `json:"subaddress"`

	// Nettype Specifies which of the three Monero networks (mainnet, stagenet, and testnet) the address belongs to.
	Nettype string `json:"nettype"`

	// OpenaliasAddress True if the address is OpenAlias-formatted.
	OpenaliasAddress bool `json:"openalias_address"`
}

// ValidateAddress Analyzes a string to determine whether it is a valid monero wallet address and returns the result and the address specifications.
func (c *Client) ValidateAddress(req *ValidateAddressRequest) (*ValidateAddressResponse, error) {
	resp := &ValidateAddressResponse{}
	err := c.Do("validate_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
