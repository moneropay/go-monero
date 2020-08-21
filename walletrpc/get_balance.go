package walletrpc

type GetBalanceRequest struct {
	// AccountIndex Return balance for this account.
	AccountIndex uint64 `json:"account_index"`

	// AddressIndices (Optional) Return balance detail for those subaddresses.
	AddressIndices []uint64 `json:"address_indices"`
}

type GetBalanceResponse struct {
	// Balance The total balance of the current monero-wallet-rpc in session.
	Balance uint64 `json:"balance"`

	// UnlockedBalance Unlocked funds are those funds that are sufficiently deep enough in the Monero blockchain to be considered safe to spend.
	UnlockedBalance uint64 `json:"unlocked_balance"`

	// MultisigImportNeeded True if importing multisig data is needed for returning a correct balance.
	MultisigImportNeeded bool `json:"multisig_import_needed"`

	// PerSubaddress Balance information for each subaddress in an account.
	PerSubaddress []Address `json:"per_subaddress"`
}

// GetBalance Return the wallet's balance.
func (c *Client) GetBalance(req *GetBalanceRequest) (*GetBalanceResponse, error) {
	resp := &GetBalanceResponse{}
	err := c.do("get_balance", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
