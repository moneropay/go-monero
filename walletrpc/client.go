package walletrpc

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

// New returns a new monero-wallet-rpc client.
func New(cfg Config) *Client {
	cl := &Client{
		addr:    cfg.Address,
		headers: cfg.CustomHeaders,
	}
	if cfg.Transport == nil {
		cl.httpcl = http.DefaultClient
	} else {
		cl.httpcl = &http.Client{
			Transport: cfg.Transport,
		}
	}
	return cl
}

type Client struct {
	httpcl  *http.Client
	addr    string
	headers map[string]string
}

func (c *Client) Do(method string, in, out interface{}) error {
	payload, err := json2.EncodeClientRequest(method, in)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.addr, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.headers != nil {
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.httpcl.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	// in theory this is only done to catch
	// any monero related errors if
	// we are not expecting any data back
	if out == nil {
		v := &json2.EmptyResponse{}
		return json2.DecodeClientResponse(resp.Body, v)
	}
	return json2.DecodeClientResponse(resp.Body, out)
}

// CreateAddress returns new address
func (c *Client) CreateAddress(req CreateAddressRequest) (*CreateAddressResponse, error) {
	resp := &CreateAddressResponse{}
	err := c.Do("create_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateAccount returns new account
func (c *Client) CreateAccount() (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := c.Do("create_account", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetBalance returns the wallet's balance.
func (c *Client) GetBalance(req GetBalanceRequest) (*GetBalanceResponse, error) {
	resp := &GetBalanceResponse{}
	err := c.Do("get_balance", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAddress returns the wallet's address.
// address - string; The 95-character hex address string of the monero-wallet-rpc in session.
func (c *Client) GetAddress(req GetAddressRequest) (*GetAddressResponse, error) {
	resp := &GetAddressResponse{}
	err := c.Do("get_address", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetHeight returns the wallet's current block height.
// height - unsigned int; The current monero-wallet-rpc's blockchain height.
// If the wallet has been offline for a long time, it may need to catch up with the daemon.
func (c *Client) GetHeight() (uint64, error) {
	jd := struct {
		Height uint64 `json:"height"`
	}{}
	err := c.Do("get_height", nil, &jd)
	if err != nil {
		return 0, err
	}
	return jd.Height, nil
}

// Transfer sends monero to a number of recipients.
func (c *Client) Transfer(req TransferRequest) (*TransferResponse, error) {
	resp := &TransferResponse{}
	err := c.Do("transfer", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TransferSplit is same as transfer, but can split into more than one tx if necessary.
func (c *Client) TransferSplit(req TransferRequest) (resp *TransferSplitResponse, err error) {
	resp = &TransferSplitResponse{}
	err = c.Do("transfer_split", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SweepDust sends all dust outputs back to the wallet's, to make them easier to spend (and mix).
func (c *Client) SweepDust() (txHashList []string, err error) {
	jd := struct {
		TxHashList []string `json:"tx_hash_list"`
	}{}
	err = c.Do("sweep_dust", nil, &jd)
	if err != nil {
		return nil, err
	}
	return jd.TxHashList, nil
}

// SweepAll sends all unlocked balance to an address.
func (c *Client) SweepAll(req SweepAllRequest) (resp *SweepAllResponse, err error) {
	resp = &SweepAllResponse{}
	err = c.Do("sweep_all", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SweepSingle sends all of a specific unlocked output to an address.
func (c *Client) SweepSingle(req SweepSingleRequest) (resp *SweepSingleResponse, err error) {
	resp = &SweepSingleResponse{}
	err = c.Do("sweep_single", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Store saves the blockchain.
func (c *Client) Store() error {
	return c.Do("store", nil, nil)
}

// GetPayments returns a list of incoming payments using a given payment id.
func (c *Client) GetPayments(paymentid string) ([]Payment, error) {
	jin := struct {
		PaymentID string `json:"payment_id"`
	}{
		paymentid,
	}
	jd := struct {
		Payments []Payment `json:"payments"`
	}{}
	err := c.Do("get_payments", &jin, &jd)
	if err != nil {
		return nil, err
	}
	return jd.Payments, nil
}

// GetBulkPayments returns a list of incoming payments using a given payment id, or a list of
// payments ids, from a given height. This method is the preferred method
// over get_payments because it has the same functionality but is more extendable.
// Either is fine for looking up transactions by a single payment ID.
//	payment_ids - array of: string
//	min_block_height - unsigned int; The block height at which to start looking for payments.
func (c *Client) GetBulkPayments(paymentids []string, minblockheight uint64) ([]Payment, error) {
	jin := struct {
		PaymentIDs     []string `json:"payment_ids"`
		MinBlockHeight uint64   `json:"min_block_height"`
	}{
		paymentids,
		minblockheight,
	}
	jd := struct {
		Payments []Payment `json:"payments"`
	}{}
	err := c.Do("get_bulk_payments", &jin, &jd)
	if err != nil {
		return nil, err
	}
	return jd.Payments, nil
}

// GetTransfers returns a list of transfers.
func (c *Client) GetTransfers(req GetTransfersRequest) (*GetTransfersResponse, error) {
	resp := &GetTransfersResponse{}
	err := c.Do("get_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTransferByTxID shows information about a transfer to/from this address.
func (c *Client) GetTransferByTxID(req GetTransferByTxidRequest) (*GetTransferByTxidResponse, error) {
	resp := &GetTransferByTxidResponse{}
	err := c.Do("get_transfer_by_txid", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// IncomingTransfers returns a list of incoming transfers to the wallet.
func (c *Client) IncomingTransfers(req GetIncomingTransferRequest) (*GetIncomingTransferResponse, error) {
	resp := &GetIncomingTransferResponse{}
	err := c.Do("incoming_transfers", &req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// QueryKey returns the spend or view private key (or mnemonic seed).
func (c *Client) QueryKey(keytype QueryKeyType) (key string, err error) {
	jin := struct {
		KeyType QueryKeyType `json:"key_type"`
	}{
		keytype,
	}
	jd := struct {
		Key string `json:"key"`
	}{}
	err = c.Do("query_key", &jin, &jd)
	if err != nil {
		return
	}
	key = jd.Key
	return
}

// MakeIntegratedAddress makes an integrated address from the wallet address and a payment id.
// payment_id - string; hex encoded; can be empty, in which case a random payment id is generated
func (c *Client) MakeIntegratedAddress(paymentid string) (integratedaddr string, err error) {
	jin := struct {
		PaymentID string `json:"payment_id"`
	}{
		paymentid,
	}
	jd := struct {
		Address string `json:"integrated_address"`
	}{}
	err = c.Do("make_integrated_address", &jin, &jd)
	if err != nil {
		return
	}
	integratedaddr = jd.Address
	return
}

// SplitIntegratedAddress retrieves the standard address and payment id corresponding to an integrated address.
func (c *Client) SplitIntegratedAddress(integratedaddr string) (paymentid, address string, err error) {
	jin := struct {
		IntegratedAddress string `json:"integrated_address"`
	}{
		integratedaddr,
	}
	jd := struct {
		Address   string `json:"standard_address"`
		PaymentID string `json:"payment_id"`
	}{}
	err = c.Do("split_integrated_address", &jin, &jd)
	if err != nil {
		return
	}
	paymentid = jd.PaymentID
	address = jd.Address
	return
}

// StopWallet stops the wallet, storing the current state.
func (c *Client) StopWallet() error {
	return c.Do("stop_wallet", nil, nil)
}

// MakeURI creates a payment URI using the official URI spec.
func (c *Client) MakeURI(req URIDef) (uri string, err error) {
	jd := struct {
		URI string `json:"uri"`
	}{}
	err = c.Do("make_uri", &req, &jd)
	if err != nil {
		return
	}
	uri = jd.URI
	return
}

// ParseURI parses a payment URI to get payment information.
func (c *Client) ParseURI(uri string) (parsed *URIDef, err error) {
	jin := struct {
		URI string `json:"uri"`
	}{
		uri,
	}
	parsed = &URIDef{}
	err = c.Do("parse_uri", &jin, parsed)
	if err != nil {
		return nil, err
	}
	return
}

// RescanBlockchain rescans blockchain from scratch.
func (c *Client) RescanBlockchain() error {
	return c.Do("rescan_blockchain", nil, nil)
}

// SetTxNotes sets arbitrary string notes for transactions.
func (c *Client) SetTxNotes(txids, notes []string) error {
	jin := struct {
		TxIDs []string `json:"txids"`
		Notes []string `json:"notes"`
	}{
		txids,
		notes,
	}
	return c.Do("set_tx_notes", &jin, nil)
}

// GetTxNotes gets string notes for transactions.
func (c *Client) GetTxNotes(txids []string) (notes []string, err error) {
	jin := struct {
		TxIDs []string `json:"txids"`
	}{
		txids,
	}
	jd := struct {
		Notes []string `json:"notes"`
	}{}
	err = c.Do("get_tx_notes", &jin, &jd)
	if err != nil {
		return nil, err
	}
	notes = jd.Notes
	return
}

// SetAttribute sets an attribute value
func (c *Client) SetAttribute(key string, value string) error {
	jin := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		key,
		value,
	}
	return c.Do("set_attribute", &jin, nil)
}

// GetAttribute gets an attribute value
func (c *Client) GetAttribute(key string) (value string, err error) {
	jin := struct {
		Key string `json:"key"`
	}{
		key,
	}
	jd := struct {
		Value string `json:"value"`
	}{}
	err = c.Do("get_attribute", &jin, &jd)
	if err != nil {
		return "", err
	}
	value = jd.Value
	return
}

// Sign signs a string.
func (c *Client) Sign(data string) (signature string, err error) {
	jin := struct {
		Data string `json:"data"`
	}{
		data,
	}
	jd := struct {
		Signature string `json:"signature"`
	}{}
	err = c.Do("sign", &jin, &jd)
	if err != nil {
		return "", err
	}
	signature = jd.Signature
	return
}

// Verify verifies a signature on a string.
func (c *Client) Verify(data, address, signature string) (good bool, err error) {
	jin := struct {
		Data      string `json:"data"`
		Address   string `json:"address"`
		Signature string `json:"signature"`
	}{
		data,
		address,
		signature,
	}
	jd := struct {
		Good bool `json:"good"`
	}{}
	err = c.Do("verify", &jin, &jd)
	if err != nil {
		return false, err
	}
	good = jd.Good
	return
}

// ExportKeyImages exports a signed set of key images.
func (c *Client) ExportKeyImages() (signedkeyimages []SignedKeyImage, err error) {
	jd := struct {
		SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
	}{}
	err = c.Do("export_key_images", nil, &jd)
	signedkeyimages = jd.SignedKeyImages
	return
}

// ImportKeyImages imports signed key images list and verify their spent status.
func (c *Client) ImportKeyImages(signedkeyimages []SignedKeyImage) (resp *ImportKeyImageResponse, err error) {
	jin := struct {
		SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
	}{
		signedkeyimages,
	}
	resp = &ImportKeyImageResponse{}
	err = c.Do("import_key_images", &jin, resp)
	if err != nil {
		return nil, err
	}
	return
}

// GetAddressBook retrieves entries from the address book.
// indexes - array of unsigned int; indices of the requested address book entries
func (c *Client) GetAddressBook(indexes []uint64) (entries []AddressBookEntry, err error) {
	jin := struct {
		Indexes []uint64 `json:"entries"`
	}{
		indexes,
	}
	jd := struct {
		Entries []AddressBookEntry `json:"entries"`
	}{}
	err = c.Do("get_address_book", &jin, &jd)
	if err != nil {
		return nil, err
	}
	entries = jd.Entries
	return
}

// AddAddressBook adds an entry to the address book.
func (c *Client) AddAddressBook(entry AddressBookEntry) (index uint64, err error) {
	entry.Index = 0
	jd := struct {
		Index uint64 `json:"index"`
	}{}
	err = c.Do("add_address_book", &entry, &jd)
	if err != nil {
		return 0, err
	}
	index = jd.Index
	return
}

// DeleteAddressBook deletes an entry from the address book.
func (c *Client) DeleteAddressBook(index uint64) error {
	jin := struct {
		Index uint64 `json:"index"`
	}{
		index,
	}
	return c.Do("delete_address_book", &jin, nil)
}

// RescanSpent rescans the blockchain for spent outputs.
func (c *Client) RescanSpent() error {
	return c.Do("rescan_spent", nil, nil)
}

// StartMining starts mining in the Monero daemon.
//	threads_count - unsigned int; Number of threads created for mining
//	do_background_mining - boolean;
//	ignore_battery - boolean;
func (c *Client) StartMining(threads uint, background, ignorebattery bool) error {
	jin := struct {
		Threads       uint `json:"threads_count"`
		Background    bool `json:"do_background_mining"`
		IgnoreBattery bool `json:"ignore_battery"`
	}{
		threads,
		background,
		ignorebattery,
	}
	return c.Do("start_mining", &jin, nil)
}

// StopMining stops mining in the Monero daemon.
func (c *Client) StopMining() error {
	return c.Do("stop_mining", nil, nil)
}

// GetLanguages gets a list of available languages for your wallet's seed.
func (c *Client) GetLanguages() (languages []string, err error) {
	jd := struct {
		Languages []string `json:"languages"`
	}{}
	err = c.Do("get_languages", nil, &jd)
	if err != nil {
		return nil, err
	}
	languages = jd.Languages
	return
}

// CreateWallet creates a new wallet. You need to have set the argument "–wallet-dir" when
// launching monero-wallet-rpc to make this work.
//   filename - string;
//    password - string;
//    language - string; Language for your wallets' seed.
func (c *Client) CreateWallet(filename, password, language string) error {
	jin := struct {
		Filename string `json:"filename"`
		Password string `json:"password"`
		Language string `json:"language"`
	}{
		filename,
		password,
		language,
	}
	return c.Do("create_wallet", &jin, nil)
}

// OpenWallet opens a wallet. You need to have set the argument "–wallet-dir" when
// launching monero-wallet-rpc to make this work.
func (c *Client) OpenWallet(filename, password string) error {
	jin := struct {
		Filename string `json:"filename"`
		Password string `json:"password"`
	}{
		filename,
		password,
	}
	return c.Do("open_wallet", &jin, nil)
}
