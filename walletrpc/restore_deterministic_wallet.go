package walletrpc

type RestoreDeterministicWalletRequest struct {
	// Name of the wallet.
	Name string `json:"name"`

	// Password of the wallet.
	Password string `json:"password"`

	// Seed Mnemonic phrase of the wallet to restore.
	Seed string `json:"seed"`

	// RestoreHeight (Optional) Block height to restore the wallet from (default = 0).
	RestoreHeight int64 `json:"restore_height"`

	// Language (Optional) Language of the mnemonic phrase in case the old language is invalid.
	Language string `json:"language"`

	// SeedOffset (Optional) Offset used to derive a new seed from the given mnemonic to recover a secret wallet from the mnemonic phrase.
	SeedOffset string `json:"seed_offset"`

	// AutosaveCurrent Whether to save the currently open RPC wallet before closing it (Defaults to true).
	AutosaveCurrent bool `json:"autosave_current"`
}

type RestoreDeterministicWalletResponse struct {
	// Address 95-character hexadecimal address of the restored wallet as a string.
	Address string `json:"address"`

	// Info Message describing the success or failure of the attempt to restore the wallet.
	Info string `json:"info"`

	// Seed Mnemonic phrase of the restored wallet, which is updated if the wallet was restored from a deprecated-style mnemonic phrase.
	Seed string `json:"seed"`

	// WasDeprecated Indicates if the restored wallet was created from a deprecated-style mnemonic phrase.
	WasDeprecated bool `json:"was_deprecated"`
}

// RestoreDeterministicWallet Create and open a wallet on the RPC server from an existing mnemonic phrase and close the currently open wallet.
func (c *Client) RestoreDeterministicWallet(req *RestoreDeterministicWalletRequest) (*RestoreDeterministicWalletResponse, error) {
	resp := &RestoreDeterministicWalletResponse{}
	err := c.Do("restore_deterministic_wallet", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
