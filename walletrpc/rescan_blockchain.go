package walletrpc

// RescanBlockchain Rescan the blockchain from scratch, losing any information which can not be recovered from the blockchain itself.
func (c *Client) RescanBlockchain() error {
	err := c.do("rescan_blockchain", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
