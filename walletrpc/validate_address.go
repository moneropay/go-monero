package walletrpc

import "context"

type ValidateAddressRequest struct {
	// The address to validate.
	Address string `json:"address"`

	// (Optional) If true, consider addresses belonging to any of the three Monero networks (mainnet, stagenet, and testnet) valid. Otherwise, only consider an address valid if it belongs to the network on which the rpc-wallet's current daemon is running (Defaults to false).
	AnyNetType bool `json:"any_net_type,omitempty"`

	// (Optional) If true, consider OpenAlias-formatted addresses valid (Defaults to false).
	AllowOpenalias bool `json:"allow_openalias,omitempty"`
}

type ValidateAddressResponse struct {
	// True if the input address is a valid Monero address.
	Valid bool `json:"valid"`

	// True if the given address is an integrated address.
	Integrated bool `json:"integrated"`

	// True if the given address is a subaddress
	Subaddress bool `json:"subaddress"`

	// Specifies which of the three Monero networks (mainnet, stagenet, and testnet) the address belongs to.
	Nettype string `json:"nettype"`

	// True if the address is OpenAlias-formatted.
	OpenaliasAddress bool `json:"openalias_address"`
}

// Analyzes a string to determine whether it is a valid monero wallet address and returns the result and the address specifications.
func (c *Client) ValidateAddress(ctx context.Context, req *ValidateAddressRequest) (*ValidateAddressResponse, error) {
	resp := &ValidateAddressResponse{}
	err := c.Do(ctx, "validate_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
