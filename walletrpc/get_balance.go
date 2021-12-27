package walletrpc

import "context"

type GetBalanceRequest struct {
	// Return balance for this account.
	AccountIndex uint64 `json:"account_index"`

	// (Optional) Return balance detail for those subaddresses.
	AddressIndices []uint64 `json:"address_indices,omitempty"`
}

type GetBalanceResponse struct {
	// The total balance of the current monero-wallet-rpc in session.
	Balance uint64 `json:"balance"`

	// Unlocked funds are those funds that are sufficiently deep enough in the Monero blockchain to be considered safe to spend.
	UnlockedBalance uint64 `json:"unlocked_balance"`

	// True if importing multisig data is needed for returning a correct balance.
	MultisigImportNeeded bool `json:"multisig_import_needed"`

	// Balance information for each subaddress in an account.
	PerSubaddress []Address `json:"per_subaddress"`
}

// Return the wallet's balance.
func (c *Client) GetBalance(ctx context.Context, req *GetBalanceRequest) (*GetBalanceResponse, error) {
	resp := &GetBalanceResponse{}
	err := c.Do(ctx, "get_balance", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
