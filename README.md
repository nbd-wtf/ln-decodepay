decodepay
---------

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/fiatjaf/ln-decodepay/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/fiatjaf/ln-decodepay/zpay32)

Simple Lightning Network BOLT11 invoice decoder with outputs similar to [c-lightning](https://github.com/ElementsProject/lightning/blob/master/doc/lightning-decodepay.7.txt) using code from [lnd](https://github.com/lightningnetwork/lnd).

This is necessary because c-lightning is a pain to deal with if you're using `description_hash` and you may want a similar API.

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
	bolt11, _ := decodepay.Decodepay("lnbcrt1231230p1pwmtcwvpp56r8664wz9eer8s7gpw07084035gj2p40g0jxumq0ywhlafme8r8qdq8dahhqucxqyjw5qrzjqwjnq83cw0t35cfcepdvlmfvfw7ref949ds8avmd78gq3j7g8kza2pgvyuqqp9gqqyqqqqqqqqqqpjqqpynp4qg4kedljtpxr48exzd38vjd5qxh8gyhlkqg2gh56dyn90smf4sj77dgupy2xadwrcf3sw5u94s08wumvhvcx5fmyp807ekd5fmsrlfflyu56r3zp47jvn2gnjkw48qxr24gp8n4r5tkcr3xa7vmtv4gxh8fsqwuz53j")
	j, _ := json.MarshalIndent(bolt11, "", "  ")
	fmt.Println(string(j))

	bolt11, _ = decodepay.Decodepay("lnbc6540n1pwap9atpp52jwdhxg3pz89e8qh26dxpjfqz5nppak70xlhqmqks4jml0tckxashp5sm6h5lymne3d90kdy3pml9us0pr2kw4zktjgyps3h34hhl0tkv7sxqrrssnp4qdkuuuwgkqyk9ltmu8jjc297j3d5tfrw4pvvacwg7hdwqdwszavlw0gga08t3x85udljaqphq29lzz0me5lpcs6rrcxuee2nezrgyny7hyxktjle6ygvrzxffem2hd7e9qj2c2tpyxlcsg6w9skguxatdyxqpk6ru20")
	j, _ = json.MarshalIndent(bolt11, "", "  ")
	fmt.Println(string(j))
}
```

outputs

```json
{
  "currency": "bcrt",
  "created_at": 1572200908,
  "expiry": 604800,
  "payee": "022b6cb7f2584c3a9f2613627649b401ae7412ffb010a45e9a692657c369ac25ef",
  "msatoshi": 123123,
  "description": "oops",
  "payment_hash": "d0cfad55c22e7233c3c80b9fe79eaf8d112506af43e46e6c0f23affea77938ce",
  "min_final_cltv_expiry": 9,
  "routes": [
    [
      {
        "pubkey": "03a5301e3873d71a6138c85acfed2c4bbc3ca4b52b607eb36df1d008cbc83d85d5",
        "short_channel_id": "330791x149x1",
        "fee_base_msat": 0,
        "fee_proportional_millionths": 200,
        "cltv_expiry_delta": 9
      }
    ]
  ]
}
{
  "currency": "bc",
  "created_at": 1573951403,
  "expiry": 3600,
  "payee": "036dce71c8b00962fd7be1e52c28be945b45a46ea858cee1c8f5dae035d01759f7",
  "msatoshi": 654000,
  "description_hash": "86f57a7c9b9e62d2becd2443bf97907846ab3aa2b2e4820611bc6b7bfdebb33d",
  "payment_hash": "549cdb9911088e5c9c17569a60c920152610f6de79bf706c168565bfbd78b1bb",
  "min_final_cltv_expiry": 9
}
```
