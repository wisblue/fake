package fake

import (
	"reflect"
	"testing"
)

func TestFillStruct(t *testing.T) {
	type A struct {
		UserName       string
		EmailAddress   string
		Place          string `fake:"StreetAddress"`
		Phone          string
		PostCode       string `fake:"zip"`
		CreditCardType string
		CreditCardNum  string
		Currency       string
		CurrencyCode   string
		Day            int
		WeekDay        string
		WeekDayShort   string
		WeekdayNum     int
		Month          string
		MonthShort     string
		MonthNum       int
		Year           int
		Longitude      float32
		Latitude       float32
		Password       string
		Color          string
		HexColor       string
		DomainName     string
		IPv4           string
		IPv6           string
		UserAgent      string
		Company        string
		JobTitle       string
		Industry       string
		Title          string
		Sentence       string
		Sentences      string
		Paragraphs     string
		FullName       string
		Gender         string
		Language       string
		Brand          string
		ProductName    string
		Product        string
	}
	a := &A{}

	for _, lang := range availLangs {
		SetLang(lang)

		a = FillStruct(a).(*A)
		//t.Logf("%+v\n", *a)

		vv := reflect.ValueOf(a).Elem()
		for j := 0; j < vv.NumField(); j++ {
			f := vv.Field(j)
			n := vv.Type().Field(j).Name
			typ := f.Type().String()
			//t.Logf("Name: %s  Kind: %s  Type: %s\n", n, f.Kind(), typ)

			if typ == "string" {
				if f.String() == "" {
					t.Log("Failed field:", n)
					t.Fail()
				}
			} else if typ == "int" {
				if f.Int() < 0 {
					t.Log("Failed field:", n, f.Int())
					t.Fail()
				}
			} else if typ == "float32" {
				if f.Float() < -180 || f.Float() > 180 {
					t.Log("Failed field:", n, f.Float())
					t.Fail()
				}
			}
		}
	}
}
