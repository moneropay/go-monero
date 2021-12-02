package walletrpc

import "context"

type DescribeTransferRequest struct {
	// (Optional) A hexadecimal string representing a set of unsigned transactions (empty for multisig transactions; non-multisig signed transactions are not supported).
	UnsignedTxset string `json:"unsigned_txset,omitempty"`

	// (Optional) A hexadecimal string representing the set of signing keys used in a multisig transaction (empty for unsigned transactions; non-multisig signed transactions are not supported).
	MultisigTxset string `json:"multisig_txset,omitempty"`
}

type DescribeTransferResponse struct {
	// The description of the transfer as a list:
	Desc []Description `json:"desc"`
}

type Description struct {
	// The sum of the inputs spent by the transaction in atomic units.
	AmountIn uint64 `json:"amount_in"`

	// The sum of the outputs created by the transaction in atomic units.
	AmountOut uint64 `json:"amount_out"`

	// List of recipients:
	Recipients []Destination `json:"recipients"`

	// The public address of the recipient.
	Address string `json:"address"`

	// The amount sent to the recipient in atomic units.
	Amount uint64 `json:"amount"`

	// The address of the change recipient.
	ChangeAddress string `json:"change_address"`

	// The amount sent to the change address in atomic units.
	ChangeAmount uint64 `json:"change_amount"`

	// The fee charged for the transaction in atomic units.
	Fee uint64 `json:"fee"`

	// Payment ID for this transfer (empty if not provided.
	PaymentId string `json:"payment_id"`

	// The number of inputs in the ring (1 real output + the number of decoys from the blockchain).
	RingSize uint64 `json:"ring_size"`

	// The number of blocks before the monero can be spent (0 for no lock).
	UnlockTime uint64 `json:"unlock_time"`

	// The number of fake outputs added to single-output transactions. Fake outputs have 0 amount and are sent to a random address.
	DummyOutputs uint64 `json:"dummy_outputs"`

	// Arbitrary transaction data in hexadecimal format.
	Extra string `json:"extra"`
}

// Returns details for each transaction in an unsigned or multisig transaction set.
func (c *Client) DescribeTransfer(ctx context.Context, req *DescribeTransferRequest) (*DescribeTransferResponse, error) {
	resp := &DescribeTransferResponse{}
	err := c.Do(ctx, "describe_transfer", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
