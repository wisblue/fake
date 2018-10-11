package fake

import (
	"testing"
)

func TestCurrencies(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := Currency()
		if v == "" {
			t.Errorf("Currency failed with lang %s", lang)
		}

		v = CurrencyCode()
		if v == "" {
			t.Errorf("CurrencyCode failed with lang %s", lang)
		}

		price := Price()
		if price == 0 {
			t.Errorf("price failed with lang %s", lang)
		}

		pricef := PriceF()
		if pricef == 0 {
			t.Errorf("pricef failed with lang %s", lang)
		}

		q := Quantity()
		if q == 0 {
			t.Errorf("Quantity failed with lang %s", lang)
		}
	}
}
