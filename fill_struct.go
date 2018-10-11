package fake

import (
	"log"
	"reflect"

	"github.com/iancoleman/strcase"
)

var stringFaker = map[string]interface{}{
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
	"Description":  Paragraphs,
	"FullName":     FullName,
	"Name":         FullName,
	"Gender":       Gender,
	"Sex":          Gender,
	"Language":     Language,
	"Brand":        Brand,
	"ProductName":  ProductName,
	"Product":      Product,
	"URL":          URL,
}

var intFaker = map[string]func() int{
	"Day":        Day,
	"WeekdayNum": WeekdayNum,
	"MonthNum":   MonthNum,
	"Year": func() int {
		return Year(1950, 2020)
	},
}

var uint64Faker = map[string]func() uint64{
	"Price":    Price,
	"Quantity": Quantity,
}

var floatFaker = map[string]func() float32{
	"Latitude":  Latitude,
	"Longitude": Longitude,
}

var float64Faker = map[string]func() float64{
	"PriceF": PriceF,
}

var fakerGroup = map[reflect.Kind]interface{}{
	reflect.String:  stringFaker,
	reflect.Int:     intFaker,
	reflect.Uint64:  uint64Faker,
	reflect.Float32: floatFaker,
	reflect.Float64: float64Faker,
}

// FillStruct fills struct field with faked data.
// FillStruct get the field tag with tag key "fake" for which
// fake function to call. If "fake" tag is not found, it will
// look at the field name if matches a fake function. Otherwise
// a error is logged.
func FillStruct(a interface{}) interface{} {
	v := reflect.ValueOf(a).Elem()
	//var tag reflect.StructTag
	return fillStruct(v, "", reflect.StructTag("")).Addr().Interface()
}

func fillStruct(v reflect.Value, name string, tag reflect.StructTag) reflect.Value {
	if tag.Get("fake") == "-" {
		return v
	}

	t := v.Kind()

	if t == reflect.String {
		var fakeFn func() string
		if fakeFunc := tag.Get("fake"); fakeFunc != "" {
			if fn, ok := stringFaker[strcase.ToCamel(fakeFunc)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn == nil {
			if fn, ok := stringFaker[strcase.ToCamel(name)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn != nil {
			v.SetString(fakeFn())
		} else {
			log.Println("Do not know how for fake ", name)
		}
	} else if t == reflect.Int || t == reflect.Int64 {
		var fakeFn func() int
		if fakeFunc := tag.Get("fake"); fakeFunc != "" {
			if fn, ok := intFaker[strcase.ToCamel(fakeFunc)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn == nil {
			if fn, ok := intFaker[strcase.ToCamel(name)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn != nil {
			v.SetInt(int64(fakeFn()))
		}
	} else if t == reflect.Float32 || t == reflect.Float64 {
		var fakeFn func() float32
		if fakeFunc := tag.Get("fake"); fakeFunc != "" {
			if fn, ok := floatFaker[strcase.ToCamel(fakeFunc)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn == nil {
			if fn, ok := floatFaker[strcase.ToCamel(name)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn != nil {
			v.SetFloat(float64(fakeFn()))
		}
	} else if t == reflect.Struct {
		vv := reflect.Indirect(v)
		for j := 0; j < vv.NumField(); j++ {
			tag := vv.Type().Field(j).Tag
			name := vv.Type().Field(j).Name
			fillStruct(vv.Field(j), name, tag)
		}
	} else if t == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			fillStruct(v.Index(i), name, tag)
		}
	} else if t == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			fillStruct(v.Index(i), name, tag)
		}
	}

	return v
}
