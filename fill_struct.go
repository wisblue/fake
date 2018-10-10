package fake

import (
	"log"
	"reflect"

	"github.com/iancoleman/strcase"
)

var stringFaker = map[string]func() string{
	"UserName":       UserName,
	"EmailAddress":   EmailAddress,
	"StreetAddress":  StreetAddress,
	"Phone":          Phone,
	"Zip":            Zip,
	"PostCode":       Zip,
	"CreditCardType": CreditCardType,
	"CreditCardNum": func() string {
		return CreditCardNum("")
	},
	"Currency":     Currency,
	"CurrencyCode": CurrencyCode,
	"WeekDay":      WeekDay,
	"WeekDayShort": WeekDayShort,
	"Month":        Month,
	"MonthShort":   MonthShort,
	"Password":     SimplePassword,
	"Color":        Color,
	"HexColor":     HexColor,
	"DomainName":   DomainName,
	"IPv4":         IPv4,
	"IPv6":         IPv6,
	"UserAgent":    UserAgent,
	"Company":      Company,
	"JobTitle":     JobTitle,
	"Industry":     Industry,
	"Title":        Title,
	"Sentence":     Sentence,
	"Sentences":    Sentences,
	"Paragraphs":   Paragraphs,
	"FullName":     FullName,
	"Gender":       Gender,
	"Language":     Language,
	"Brand":        Brand,
	"ProductName":  ProductName,
	"Product":      Product,
}

var intFaker = map[string]func() int{
	"Day":        Day,
	"WeekdayNum": WeekdayNum,
	"MonthNum":   MonthNum,
	"Year": func() int {
		return Year(1950, 2020)
	},
}

var floatFaker = map[string]func() float32{
	"Latitude":  Latitude,
	"Longitude": Longitude,
}

// FillStruct fills struct field with faked data.
// FillStruct get the field tag with tag key "fake" for which
// fake function to call. If "fake" tag is not found, it will
// look at the field name if matches a fake function. Otherwise
// a error is logged.
func FillStruct(a interface{}) interface{} {
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		n := v.Type().Field(j).Name
		tag := v.Type().Field(j).Tag
		t := f.Type().String()

		if f.CanSet() == false {
			continue
		}

		if t == "string" {
			var fakeFn func() string
			if fakeFunc := tag.Get("fake"); fakeFunc != "" {
				if fn, ok := stringFaker[strcase.ToCamel(fakeFunc)]; ok {
					fakeFn = fn
				}
			}
			if fakeFn == nil {
				if fn, ok := stringFaker[strcase.ToCamel(n)]; ok {
					fakeFn = fn
				}
			}
			if fakeFn != nil {
				f.SetString(fakeFn())
			} else {
				log.Println("Do not know how for fake ", n)
			}
		} else if t == "int" {
			var fakeFn func() int
			if fakeFunc := tag.Get("fake"); fakeFunc != "" {
				if fn, ok := intFaker[strcase.ToCamel(fakeFunc)]; ok {
					fakeFn = fn
				}
			}
			if fakeFn == nil {
				if fn, ok := intFaker[strcase.ToCamel(n)]; ok {
					fakeFn = fn
				}
			}
			if fakeFn != nil {
				f.SetInt(int64(fakeFn()))
			}
		} else if t == "float32" {
			var fakeFn func() float32
			if fakeFunc := tag.Get("fake"); fakeFunc != "" {
				if fn, ok := floatFaker[strcase.ToCamel(fakeFunc)]; ok {
					fakeFn = fn
				}
			}
			if fakeFn == nil {
				if fn, ok := floatFaker[strcase.ToCamel(n)]; ok {
					fakeFn = fn
				}
			}
			if fakeFn != nil {
				f.SetFloat(float64(fakeFn()))
			}
		}
	}

	return a
}
