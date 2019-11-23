package decodepay_gjson

import (
	"encoding/json"

	"github.com/fiatjaf/ln-decodepay"
	"github.com/tidwall/gjson"
)

func Decodepay(bolt11 string) (gjson.Result, error) {
	decoded, err := decodepay.Decodepay(bolt11)
	if err != nil {
		return gjson.Result{}, err
	}

	jbytes, err := json.Marshal(decoded)
	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.ParseBytes(jbytes), nil
}
