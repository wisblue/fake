package fake

import (
	"net"
	"strings"

	"github.com/corpix/uarand"
	"github.com/iancoleman/strcase"
	"github.com/mozillazg/go-pinyin"
)

type userNameFuncType struct {
	makePersonName        func(lang, gender string) string
	personNameAsUserName  func(lang, gender string) string
	randomWordsAsUserName func() string
	characterWithLastName func(lang, gender string) string
}

var userNameFuncs = map[string]userNameFuncType{
	"cn": userNameFuncType{
		makePersonName: func(lang, gender string) string {
			return lookup(lang, gender+"_last_names", false) + lookup(lang, gender+"_first_names", false)
		},
		personNameAsUserName: func(lang, gender string) string {
			name := lookup(lang, gender+"_last_names", false) + lookup(lang, gender+"_first_names", false)
			py := pinyin.LazyConvert(name, nil)
			namepy := ""
			if len(py) > 1 {
				namepy = py[0] + " " + strings.Join(py[1:], "")
			} else {
				namepy = py[0]
			}
			if r.Intn(1) == 0 {
				return strcase.ToCamel(namepy)
			}
			return strcase.ToSnake(namepy)
		},
		randomWordsAsUserName: func() string {
			n := r.Intn(2) + 1
			words := make([]string, n)
			for i := 0; i < n; i++ {
				words[i] = Word()
			}
			strs_py := []string{}
			for _, word := range words {
				py := pinyin.LazyConvert(word, nil)
				str := strings.Join(py, "")
				strs_py = append(strs_py, str)
			}
			s := strings.Join(strs_py, " ")
			return strcase.ToSnake(s)
		},
		characterWithLastName: func(lang, gender string) string {
			s := lookup(lang, gender+"_first_names", false)
			py := pinyin.LazyConvert(s, nil)
			return Character() + strings.Join(py, "")
		},
	},
	"en": userNameFuncType{
		makePersonName:        defaultMakePersonName,
		personNameAsUserName:  defaultPersonNameAsUserName,
		randomWordsAsUserName: defaultRandomWordsAsUserName,
		characterWithLastName: defaultCharacterWithLastName,
	},
	"ru": userNameFuncType{
		makePersonName:        defaultMakePersonName,
		personNameAsUserName:  defaultPersonNameAsUserName,
		randomWordsAsUserName: defaultRandomWordsAsUserName,
		characterWithLastName: defaultCharacterWithLastName,
	},
}

func defaultMakePersonName(lang, gender string) string {
	return lookup(lang, gender+"_first_names", false) + " " + lookup(lang, gender+"_last_names", false)
}

func defaultPersonNameAsUserName(lang, gender string) string {
	return lookup(lang, gender+"_last_names", false) + lookup(lang, gender+"_first_names", false)
}

func defaultRandomWordsAsUserName() string {
	return strings.Replace(WordsN(r.Intn(3)+1), " ", "_", -1)
}

func defaultCharacterWithLastName(lang, gender string) string {
	return Character() + lookup(lang, gender+"_last_names", false)
}

// UserName generates user name in one of the following forms
// first name + last name, letter + last names or concatenation of from 1 to 3 lowercased words
func UserName() string {
	gender := randGender()

	switch 1 { //r.Intn(3) {
	case 0:
		return userNameFuncs[lang].personNameAsUserName(lang, gender)
	case 1:
		return userNameFuncs[lang].characterWithLastName(lang, gender)
	default:
		return userNameFuncs[lang].randomWordsAsUserName()
	}
}

// TopLevelDomain generates random top level domain
func TopLevelDomain() string {
	return lookup(lang, "top_level_domains", true)
}

// DomainName generates random domain name
func DomainName() string {
	return strings.ToLower(lookup("en", "companies", true)) + "." + TopLevelDomain()
}

// EmailAddress generates email address
func EmailAddress() string {
	return UserName() + "@" + DomainName()
}

// EmailSubject generates random email subject
func EmailSubject() string {
	return Sentence()
}

// EmailBody generates random email body
func EmailBody() string {
	return Paragraphs()
}

// DomainZone generates random domain zone
func DomainZone() string {
	return lookup(lang, "domain_zones", true)
}

// IPv4 generates IPv4 address
func IPv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(r.Intn(256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 generates IPv6 address
func IPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(r.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// UserAgent generates a random user agent.
func UserAgent() string {
	return uarand.GetRandom()
}
