package fake

import (
	"testing"
)

func TestTime(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := UnixTime()
		if v == 0 {
			t.Errorf("RandomUnixTime failed with lang %s", lang)
		}

		vs := Timestamp()
		if vs == "" {
			t.Errorf("Timestamp failed with lang %s", lang)
		}
	}
}
