package walletrpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/rpc/v2/json2"

	"github.com/gabstv/httpdigest"
	//"github.com/stretchr/testify/assert"
)

//
//func TestClient(t *testing.T) {
//
//	testClientGetAddress(t)
//	testClientGetBalance(t)
//}
//
//func testClientGetAddress(t *testing.T) {
//	//
//	// server setup
//	sv0 := basicTestServer([]testfn{
//		func(method string, params *json.RawMessage, w http.ResponseWriter, r *http.Request) bool {
//			if method == "getaddress" {
//				r0 := struct {
//					Address string `json:"address"`
//				}{
//					"45eoXYNHC4LcL2Hh42T9FMPTmZHyDEwDbgfBEuNj3RZUek8A4og4KiCfVL6ZmvHBfCALnggWtHH7QHF8426yRayLQq7MLf5",
//				}
//				writerpcResponseOK(&r0, w)
//				return true
//			}
//			return false
//		},
//	})
//	defer sv0.Close()
//	//
//	// test starts here
//	rpccl := New(Config{
//		Address: sv0.URL + "/json_rpc",
//	})
//	addr, err := rpccl.GetAddress()
//	assert.NoError(t, err)
//	assert.Equal(t, "45eoXYNHC4LcL2Hh42T9FMPTmZHyDEwDbgfBEuNj3RZUek8A4og4KiCfVL6ZmvHBfCALnggWtHH7QHF8426yRayLQq7MLf5", addr)
//}
//
//func testClientGetBalance(t *testing.T) {
//	//
//	// server setup
//	sv0 := basicTestServer([]testfn{
//		func(method string, params *json.RawMessage, w http.ResponseWriter, r *http.Request) bool {
//			if method == "getbalance" {
//				r0 := struct {
//					Balance  uint6464 `json:"balance"`
//					Unlocked uint6464 `json:"unlocked_balance"`
//				}{
//					1e12,
//					1e13,
//				}
//				writerpcResponseOK(&r0, w)
//				return true
//			}
//			return false
//		},
//	})
//	defer sv0.Close()
//	//
//	// test starts here
//	rpccl := New(Config{
//		Address: sv0.URL + "/json_rpc",
//	})
//	balance, unlocked, err := rpccl.GetBalance()
//	assert.NoError(t, err)
//	// 1 XMR
//	assert.Equal(t, uint6464(1000000000000), balance)
//	// 10 XMR
//	assert.Equal(t, uint6464(10000000000000), unlocked)
//}

//TODO: write more server stubs
//
//

type testfn = func(method string, params *json.RawMessage, w http.ResponseWriter, r *http.Request) bool

func basicTestServer(tests []testfn) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/json_rpc" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		var c clientRequest
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		for _, v := range tests {
			if v(c.Method, c.Params, w, r) {
				return
			}
		}
		// return method not found
		writerpcResponseError(ErrUnknown, "test this in curl with the real rpc", w)
	}))
}

func writerpcResponseOK(result interface{}, w http.ResponseWriter) {
	r := &clientResponse{
		Version: "2.0",
		Result:  result,
	}
	v, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(v)
}

func writerpcResponseError(code ErrorCode, message string, w http.ResponseWriter) {
	r := &clientResponse{
		Version: "2.0",
		Result:  nil,
		Error: &WalletError{
			Code:    code,
			Message: message,
		},
	}
	v, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(v)
}

// clientRequest represents a JSON-RPC request received by the server.
type clientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`
	// A String containing the name of the method to be invoked.
	Method string `json:"method"`
	// Object to pass as request parameter to the method.
	Params *json.RawMessage `json:"params"`
	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id uint64 `json:"id"`
}

// clientResponse represents a JSON-RPC response returned to a client.
type clientResponse struct {
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

func newClient() {

	client := New(Config{
		Address: "http://127.0.0.1:18082/json_rpc",
	})

	// check wallet balance
	balance, err := client.GetBalance(&GetBalanceRequest{
		AccountIndex: 0,
	})

	// there are two types of error that can happen:
	//   connection errors
	//   monero wallet errors
	// connection errors are pretty much unicorns if everything is on the
	// same instance (unless your OS hit an open files limit or something)
	if err != nil {
		if iswerr, werr := GetWalletError(err); iswerr {
			// it is a monero wallet error
			fmt.Printf("Wallet error (id:%v) %v\n", werr.Code, werr.Message)
			os.Exit(1)
		}
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Balance:", XMRToDecimal(balance.Balance))
	fmt.Println("Unlocked balance:", XMRToDecimal(balance.UnlockedBalance))
}

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestClient_GetBulkPayments(t *testing.T) {
	skipCI(t)

	trans := httpdigest.New("username", "password")

	client := New(Config{
		Address:   "http://127.0.0.1:18085/json_rpc",
		Transport: trans,
	})

	ids := []string{
		"c7bec4ff888f29fd9d4a2f9299109799d298430a717bffe61b696b08c00ef5d9",
		"0207cac71ea09d9a5c48ce0c0656788b9358753805ff66150bc615fab4567666",
		"319bcfc82906d83b23e42878ad58a32afae31ea47fecf6740578dea24e72bf1f",
		//"a30f46f0189c2974281815f908ec91d44ca09987a0cf90211234567890abf5ac", // do not get withdrawal by payment id
	}

	payment, err := client.GetBulkPayments(&GetBulkPaymentsRequest{
		PaymentIds:     ids,
		MinBlockHeight: 1881753,
	})
	if err != nil {
		t.Log(err)
	}

	t.Log(len(payment.Payments))
	t.Logf("%v\n", payment)

}

func TestClient_Transfer(t *testing.T) {
	skipCI(t)

	trans := httpdigest.New("username", "password")

	client := New(Config{
		Address:   "http://127.0.0.1:18085/json_rpc",
		Transport: trans,
	})
	req := TransferRequest{

		Destinations: []Destination{
			{
				Amount:  100,
				Address: "41hgaCaFC9MB1myjiEH6fZ4pkRm692gkNTqq85UPs9ZaDD65pDyYfq6UxvMvwCy46bPJRJu2hZ3NS6n6znSJoFEWN1pUcjG",
			},
		},
		Mixin: 3,
		// RingSize: default RingSize = mixin + 1
		//PaymentId:     "a30f46f0189c2974281815f908ec91d44ca09987a0cf90211234567890abf5ac",
		Priority:      PriorityUnimportant,
		GetTxHex:      true,
		GetTxMetadata: false,
	}

	payload, err := json2.EncodeClientRequest("transfer", req)
	if err != nil {
		t.Logf("%v", payload)
		t.Fatalf("%v", err)
	}
	t.Logf("%v", string(payload))

	resp, err := client.Transfer(&req)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%v\n", resp)
}

func TestClient_GetTransferByTxID(t *testing.T) {
	skipCI(t)

	trans := httpdigest.New("username", "password")

	client := New(Config{
		Address:   "http://127.0.0.1:18085/json_rpc",
		Transport: trans,
	})
	resp, err := client.GetTransferByTxid(&GetTransferByTxidRequest{
		Txid: "25196f09a12ec5f5127ef0e0bba7228cbce22e885c0b959545ef65eea03ea15d",
	})
	if err != nil {
		t.Logf("%v", err)
	}

	t.Logf("resp: %v", *resp)
}
