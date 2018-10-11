package fake

import (
	"math"
)

// Currency generates currency name
func Currency() string {
	return lookup(lang, "currencies", true)
}

// CurrencyCode generates currency code
func CurrencyCode() string {
	return lookup(lang, "currency_codes", true)
}

// Price generates uint64 price in cent
func Price() uint64 {
	return uint64(r.Intn(98)+1) * uint64(math.Pow10(r.Intn(5)+2))
}

// PriceF generates float64 price with precision of 2
func PriceF() float64 {
	return float64(Price()) / 100.0
}

func Quantity() uint64 {
	return uint64(r.Intn(9)+1) * uint64(math.Pow10(r.Intn(1)))

}
