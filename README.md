decodepay
---------

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/fiatjaf/ln-decodepay/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/fiatjaf/ln-decodepay/zpay32)

Simple Lightning Network BOLT11 invoice decoder with outputs similar to [c-lightning](https://github.com/ElementsProject/lightning/blob/master/doc/lightning-decodepay.7.txt) using code from [lnd](https://github.com/lightningnetwork/lnd).

### install

```
go get -u github.com/fiatjaf/ln-decodepay
```

### use

```go
package main

import (
	"encoding/json"
	"fmt"

	decodepay "github.com/fiatjaf/ln-decodepay"
)

func main() {
	bolt11, _ := decodepay.Decodepay("lnbc230n1pw825fwpp5a8gqlm5jqfh0s4xex58n8fev3hcww88nhu2d7d2m46rxyjnlqfpsdq0wd6hqetjw3jhxaqxqr3jscqp2rzjq2j5m6udpugaglr02h8vtennqc7fltfxr92earv84cltfsupv6zyjzyd3vqqjdsqqqqqqq86qqqqqqsqjq8vy7dtmlxwx8u448gwzanpua3eu9yafq9vr3w7ls8ferxcaffc3rrma53ah86cysh5598vgdarm40306raq5cmhpkemtwjle4zllqtsqg7jldf")
	j, _ := json.MarshalIndent(bolt11, "", "  ")
	fmt.Println(string(j))
}
```

outputs

```json
{
  "currency": "bc",
  "created_at": 1551192366,
  "expiry": 18000,
  "payee": "02c16cca44562b590dd279c942200bdccfd4f990c3a69fad620c10ef2f8228eaff",
  "msatoshi": 23000,
  "description": "supertest",
  "payment_hash": "e9d00fee92026ef854d9350f33a72c8df0e71cf3bf14df355bae86624a7f0243",
  "min_final_cltv_expiry": 10
}
```
