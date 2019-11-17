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
