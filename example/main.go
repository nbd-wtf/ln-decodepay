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
