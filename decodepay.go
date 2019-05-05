package decodepay

import (
	"encoding/hex"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/lightningnetwork/lnd/zpay32"
)

var DefaultCurrency = "bc"

func Decodepay(bolt11 string) (Bolt11, error) {
	return DecodepayWithCurrency(DefaultCurrency, bolt11)
}

func DecodepayWithCurrency(currency, bolt11 string) (Bolt11, error) {
	inv, err := zpay32.Decode(bolt11, &chaincfg.Params{Bech32HRPSegwit: currency})
	if err != nil {
		return Bolt11{}, err
	}

	var msat uint64
	if inv.MilliSat != nil {
		msat = uint64(*inv.MilliSat)
	}

	var desc string
	if inv.Description != nil {
		desc = *inv.Description
	}

	return Bolt11{
		MSatoshi:           msat,
		PaymentHash:        hex.EncodeToString(inv.PaymentHash[:]),
		Description:        desc,
		Payee:              hex.EncodeToString(inv.Destination.SerializeCompressed()),
		CreatedAt:          uint64(inv.Timestamp.Unix()),
		Expiry:             uint64(inv.Expiry() / time.Second),
		MinFinalCLTVExpiry: inv.MinFinalCLTVExpiry(),
		Currency:           DefaultCurrency,
	}, nil
}

type Bolt11 struct {
	Currency           string `json:"currency"`
	CreatedAt          uint64 `json:"created_at"`
	Expiry             uint64 `json:"expiry"`
	Payee              string `json:"payee"`
	MSatoshi           uint64 `json:"msatoshi"`
	Description        string `json:"description"`
	PaymentHash        string `json:"payment_hash"`
	MinFinalCLTVExpiry uint64 `json:"min_final_cltv_expiry"`
}
