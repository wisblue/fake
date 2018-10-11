package fake

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestFillStruct(t *testing.T) {
	type A struct {
		UserName string
		Profile  struct {
			EmailAddresses [2]string `fake:"EmailAddress"`
			Place          string    `fake:"StreetAddress"`
			Phone          string
			PostCode       string `fake:"zip"`
		}
		CreditCardType string
		CreditCardNum  string
		Currency       string
		CurrencyCode   string `fake:"-"` // skip
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
		URL            string
		Description    string
		Price          uint64
		Quantiry       uint64
		PriceF         float64
	}
	a := &A{}

	for _, lang := range availLangs {
		SetLang(lang)

		a = FillStruct(a).(*A)
		//t.Logf("%+v\n", *a)

		t.Log(a.Price)

		if a.CurrencyCode != "" {
			t.Errorf("Expect CurrencyCode to be empty. got %s\n", a.CurrencyCode)
		}

		vv := reflect.ValueOf(a).Elem()
		for j := 0; j < vv.NumField(); j++ {
			f := vv.Field(j)
			n := vv.Type().Field(j).Name
			kind := f.Kind()
			//t.Logf("Name: %s  Kind: %s  Type: %s\n", n, f.Kind(), typ)

			if kind == reflect.String {
				if f.String() == "" && n != "CurrencyCode" {
					t.Log("Failed field:", n)
					t.Fail()
				}
			} else if kind == reflect.Int {
				if f.Int() < 0 {
					t.Log("Failed field:", n, f.Int())
					t.Fail()
				}
			} else if kind == reflect.Float32 {
				if f.Float() < -180 || f.Float() > 180 {
					t.Log("Failed field:", n, f.Float())
					t.Fail()
				}
			} else if kind == reflect.Struct {
				// check fields of embedded struct
				for i := 0; i < f.NumField(); i++ {
					v := f.Field(i)
					if v.Type().Kind() == reflect.Array {
						// check elements of array
						for k := 0; k < v.Len(); k++ {
							item := v.Index(k)
							if item.Kind() != reflect.String {
								t.Fail()
							}
							if item.String() == "" {
								t.Fail()
							}
						}
					} else if v.Type().Kind() == reflect.String {
						if v.String() == "" {
							t.Fail()
						}
					}
				}
			}
		}
	}
}
