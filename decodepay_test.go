package decodepay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodepay(t *testing.T) {
	t.Run("decode valid bolt11 lowercase", func(t *testing.T) {
		bolt11 := "lnbc15u1p3xnhl2pp5jptserfk3zk4qy42tlucycrfwxhydvlemu9pqr93tuzlv9cc7g3sdqsvfhkcap3xyhx7un8cqzpgxqzjcsp5f8c52y2stc300gl6s4xswtjpc37hrnnr3c9wvtgjfuvqmpm35evq9qyyssqy4lgd8tj637qcjp05rdpxxykjenthxftej7a2zzmwrmrl70fyj9hvj0rewhzj7jfyuwkwcg9g2jpwtk3wkjtwnkdks84hsnu8xps5vsq4gj5hs"

		expected := Bolt11{
			Currency:           "bc",
			CreatedAt:          1651105770,
			Expiry:             600,
			Payee:              "03d6b14390cd178d670aa2d57c93d9519feaae7d1e34264d8bbb7932d47b75a50d",
			MSatoshi:           1500000,
			Description:        "bolt11.org",
			DescriptionHash:    "",
			PaymentHash:        "90570c8d3688ad5012aa5ff982606971ae46b3f9df0a100cb15f05f61718f223",
			MinFinalCLTVExpiry: 40,
			Route:              nil,
		}

		actual, err := Decodepay(bolt11)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("decode valid bolt11 uppercase", func(t *testing.T) {
		bolt11 := "LNBC15U1P3XNHL2PP5JPTSERFK3ZK4QY42TLUCYCRFWXHYDVLEMU9PQR93TUZLV9CC7G3SDQSVFHKCAP3XYHX7UN8CQZPGXQZJCSP5F8C52Y2STC300GL6S4XSWTJPC37HRNNR3C9WVTGJFUVQMPM35EVQ9QYYSSQY4LGD8TJ637QCJP05RDPXXYKJENTHXFTEJ7A2ZZMWRMRL70FYJ9HVJ0REWHZJ7JFYUWKWCG9G2JPWTK3WKJTWNKDKS84HSNU8XPS5VSQ4GJ5HS"

		expected := Bolt11{
			Currency:           "bc",
			CreatedAt:          1651105770,
			Expiry:             600,
			Payee:              "03d6b14390cd178d670aa2d57c93d9519feaae7d1e34264d8bbb7932d47b75a50d",
			MSatoshi:           1500000,
			Description:        "bolt11.org",
			DescriptionHash:    "",
			PaymentHash:        "90570c8d3688ad5012aa5ff982606971ae46b3f9df0a100cb15f05f61718f223",
			MinFinalCLTVExpiry: 40,
			Route:              nil,
		}

		actual, err := Decodepay(bolt11)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Returns error for invalid bolt11 invoice", func(t *testing.T) {
		bolt11 := "lnbc1234Invalid"

		_, err := Decodepay(bolt11)
		assert.Error(t, err)
	})
}
