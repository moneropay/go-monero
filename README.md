# go-monero
[![GoDoc](https://godoc.org/gitlab.com/moneropay/go-monero/walletrpc?status.svg)](https://godoc.org/gitlab.com/moneropay/go-monero/walletrpc)

This package provides Golang bindings for monero-wallet-rpc calls.
It also contains functions for XMR/atomic unit conversion.

Unlike other wallet-rpc bindings, this package is actively maintained and it is complete
with all the rpc methods. While using this package for our other project [MoneroPay](https://gitlab.com/moneropay/moneropay),
we have realized that all the other forks don't handle optional parameters correctly and send them anyway,
therefore causing undefined behavior.

### Installation
```sh
go get -u gitlab.com/moneropay/go-monero/walletrpc
```

### Example
```sh
monero-wallet-rpc --detach \
	--rpc-bind-port 18083 \
	--wallet-file /home/moneropay/wallet \
	--password s3cure \
	--daemon-login kernal:s3cure \
	--rpc-login kernal:s3cure
```
```go
package main

import (
	"fmt"
	"log"

	"github.com/gabstv/httpdigest"

	"gitlab.com/moneropay/go-monero/walletrpc"
)

func main() {
	// username: kernal, password: s3cure
	client := walletrpc.New(walletrpc.Config{
		Address: "http://127.0.0.1:18083/json_rpc",
		Transport: httpdigest.New("kernal", "s3cure"), // Remove if no auth.
	})
	resp, err := client.GetBalance(&walletrpc.GetBalanceRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total balance:", walletrpc.XMRToDecimal(resp.Balance))
	fmt.Println("Unlocked balance:", walletrpc.XMRToDecimal(resp.UnlockedBalance))
}
```

### Special Thanks
This version of `go-monero` is based on the works of:
- [LuaxY](https://github.com/LuaxY/go-monero)
- [gabstv](https://github.com/gabstv/go-monero)
- [carryon](https://github.com/carryon/go-monero)
- [TemaKozyrev](https://github.com/TemaKozyrev/go-monero)
- [ibclabs](https://github.com/ibclabs/go-monero)
- [Fraudmarc](https://github.com/Fraudmarc/go-monero)
- [leonklingele](https://github.com/leonklingele/go-moner)
