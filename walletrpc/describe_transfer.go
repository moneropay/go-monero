package walletrpc

type DescribeTransferRequest struct {
	// UnsignedTxset (Optional) A hexadecimal string representing a set of unsigned transactions (empty for multisig transactions; non-multisig signed transactions are not supported).
	UnsignedTxset string `json:"unsigned_txset"`

	// MultisigTxset (Optional) A hexadecimal string representing the set of signing keys used in a multisig transaction (empty for unsigned transactions; non-multisig signed transactions are not supported).
	MultisigTxset string `json:"multisig_txset"`
}

type DescribeTransferResponse struct {
	// Desc The description of the transfer as a list:
	Desc []Description `json:"desc"`
}

type Description struct {
	// AmountIn (64 bit); The sum of the inputs spent by the transaction in atomic units.
	AmountIn uint64 `json:"amount_in"`

	// AmountOut (64 bit); The sum of the outputs created by the transaction in atomic units.
	AmountOut uint64 `json:"amount_out"`

	// Recipients list of recipients:
	Recipients []Destination `json:"recipients"`

	// Address The public address of the recipient.
	Address string `json:"address"`

	// Amount The amount sent to the recipient in atomic units.
	Amount uint64 `json:"amount"`

	// ChangeAddress The address of the change recipient.
	ChangeAddress string `json:"change_address"`

	// ChangeAmount The amount sent to the change address in atomic units.
	ChangeAmount uint64 `json:"change_amount"`

	// Fee The fee charged for the transaction in atomic units.
	Fee uint64 `json:"fee"`

	// PaymentId payment ID for this transfer (empty if not provided.
	PaymentId string `json:"payment_id"`

	// RingSize The number of inputs in the ring (1 real output + the number of decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`

	// UnlockTime The number of blocks before the monero can be spent (0 for no lock).
	UnlockTime uint64 `json:"unlock_time"`

	// DummyOutputs The number of fake outputs added to single-output transactions. Fake outputs have 0 amount and are sent to a random address.
	DummyOutputs uint64 `json:"dummy_outputs"`

	// Extra Arbitrary transaction data in hexadecimal format.
	Extra string `json:"extra"`
}

// DescribeTransfer Returns details for each transaction in an unsigned or multisig transaction set.
func (c *Client) DescribeTransfer(req *DescribeTransferRequest) (*DescribeTransferResponse, error) {
	resp := &DescribeTransferResponse{}
	err := c.Do("describe_transfer", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
