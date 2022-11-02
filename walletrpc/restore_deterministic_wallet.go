package walletrpc

import "context"

type RestoreDeterministicWalletRequest struct {
	// Name of the wallet.
	Name string `json:"name"`

	// Password of the wallet.
	Password string `json:"password"`

	// Mnemonic phrase of the wallet to restore.
	Seed string `json:"seed"`

	// (Optional) Block height to restore the wallet from (default = 0).
	RestoreHeight uint64 `json:"restore_height,omitempty"`

	// (Optional) Language of the mnemonic phrase in case the old language is invalid.
	Language string `json:"language,omitempty"`

	// (Optional) Offset used to derive a new seed from the given mnemonic to recover a secret wallet from the mnemonic phrase.
	SeedOffset string `json:"seed_offset,omitempty"`

	// Whether to save the currently open RPC wallet before closing it (Defaults to true).
	AutosaveCurrent bool `json:"autosave_current"`
}

type RestoreDeterministicWalletResponse struct {
	// 95-character hexadecimal address of the restored wallet as a string.
	Address string `json:"address"`

	// Message describing the success or failure of the attempt to restore the wallet.
	Info string `json:"info"`

	// Mnemonic phrase of the restored wallet, which is updated if the wallet was restored from a deprecated-style mnemonic phrase.
	Seed string `json:"seed"`

	// Indicates if the restored wallet was created from a deprecated-style mnemonic phrase.
	WasDeprecated bool `json:"was_deprecated"`
}

// Create and open a wallet on the RPC server from an existing mnemonic phrase and close the currently open wallet.
func (c *Client) RestoreDeterministicWallet(ctx context.Context, req *RestoreDeterministicWalletRequest) (*RestoreDeterministicWalletResponse, error) {
	resp := &RestoreDeterministicWalletResponse{}
	err := c.Do(ctx, "restore_deterministic_wallet", &req, resp)
	return resp, err
}
