package decodepay

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/lightningnetwork/lnd/zpay32"
)

func ChainFromCurrency(currency string) *chaincfg.Params {
	if strings.HasPrefix(currency, "bcrt") {
		return &chaincfg.RegressionNetParams
	} else if strings.HasPrefix(currency, "tb") {
		return &chaincfg.TestNet3Params
	} else if strings.HasPrefix(currency, "sb") {
		return &chaincfg.SimNetParams
	} else {
		return &chaincfg.MainNetParams
	}
}

func Decodepay(bolt11 string) (Bolt11, error) {
	return DecodepayWithChain(
		ChainFromCurrency(bolt11[2:]),
		bolt11,
	)
}

func DecodepayWithChain(chain *chaincfg.Params, bolt11 string) (Bolt11, error) {
	inv, err := zpay32.Decode(bolt11, chain)
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

	var deschash string
	if inv.DescriptionHash != nil {
		dh := *inv.DescriptionHash
		deschash = hex.EncodeToString(dh[:])
	}

	var routes [][]Hop
	if inv.RouteHints != nil {
		routes = make([][]Hop, len(inv.RouteHints))
		for ri, r := range inv.RouteHints {
			route := make([]Hop, len(r))
			for hi, h := range r {
				scid := h.ChannelID

				route[hi] = Hop{
					PubKey: hex.EncodeToString(h.NodeID.SerializeCompressed()),
					ShortChannelId: fmt.Sprintf("%dx%dx%d",
						scid>>40&0xFFFFFF, scid>>16&0xFFFFFF, scid&0xFFFF),
					FeeBaseMsat:               int(h.FeeBaseMSat),
					FeeProportionalMillionths: int(h.FeeProportionalMillionths),
					CLTVExpiryDelta:           int(h.CLTVExpiryDelta),
				}
			}
			routes[ri] = route
		}
	}

	return Bolt11{
		MSatoshi:           msat,
		PaymentHash:        hex.EncodeToString(inv.PaymentHash[:]),
		Description:        desc,
		DescriptionHash:    deschash,
		Payee:              hex.EncodeToString(inv.Destination.SerializeCompressed()),
		CreatedAt:          uint64(inv.Timestamp.Unix()),
		Expiry:             uint64(inv.Expiry() / time.Second),
		MinFinalCLTVExpiry: inv.MinFinalCLTVExpiry(),
		Currency:           inv.Net.Bech32HRPSegwit,
		Route:              routes,
	}, nil
}

type Bolt11 struct {
	Currency           string  `json:"currency"`
	CreatedAt          uint64  `json:"created_at"`
	Expiry             uint64  `json:"expiry"`
	Payee              string  `json:"payee"`
	MSatoshi           uint64  `json:"msatoshi"`
	Description        string  `json:"description,omitempty"`
	DescriptionHash    string  `json:"description_hash,omitempty"`
	PaymentHash        string  `json:"payment_hash"`
	MinFinalCLTVExpiry uint64  `json:"min_final_cltv_expiry"`
	Route              [][]Hop `json:"routes,omitempty"`
}

type Hop struct {
	PubKey                    string `json:"pubkey"`
	ShortChannelId            string `json:"short_channel_id"`
	FeeBaseMsat               int    `json:"fee_base_msat"`
	FeeProportionalMillionths int    `json:"fee_proportional_millionths"`
	CLTVExpiryDelta           int    `json:"cltv_expiry_delta"`
}
