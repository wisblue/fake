package fake

import (
	"strings"
)

// Brand generates brand name
func Brand() string {
	cmp := Company()
	if lang == "cn" {
		s := strings.Trim(cmp, "有限公司")
		s = strings.Trim(s, "控股")
		s = strings.Trim(s, "有限责任")
		s = strings.Trim(s, "股份")
		return s
	}
	return cmp
}

// ProductName generates product name
func ProductName() string {
	productName := lookup(lang, "adjectives", true) + localeFormat[lang].wordDelimiter + lookup(lang, "nouns", true)
	if r.Intn(2) == 1 {
		productName = lookup(lang, "adjectives", true) + localeFormat[lang].wordDelimiter + productName
	}
	return productName
}

// Product generates product title as brand + product name
func Product() string {
	return Brand() + localeFormat[lang].wordDelimiter + ProductName()
}

// Model generates model name that consists of letters and digits, optionally with a hyphen between them
func Model() string {
	seps := []string{"", " ", "-"}
	return CharactersN(r.Intn(3)+1) + seps[r.Intn(len(seps))] + Digits()
}
