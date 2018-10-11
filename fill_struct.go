package fake

import (
	"log"
	"reflect"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

var fakerFuncs = map[string]interface{}{
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
	"Day":          Day,
	"WeekdayNum":   WeekdayNum,
	"MonthNum":     MonthNum,
	"Year": func() int {
		return Year(1950, 2020)
	},
	"Price":     Price,
	"Quantity":  Quantity,
	"Latitude":  Latitude,
	"Longitude": Longitude,
	"PriceF":    PriceF,
	"CreatedAt": Time,
	"KSUID":     KSUID,
	"UUID":      UUID,
	"OpenID":    OpenID,
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

	if t == reflect.Struct &&
		v.Type().String() != "time.Time" {
		vv := reflect.Indirect(v)
		for j := 0; j < vv.NumField(); j++ {
			tag := vv.Type().Field(j).Tag
			name := vv.Type().Field(j).Name
			fillStruct(vv.Field(j), name, tag)
		}
	} else if t == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			fillStruct(v.Index(i), inflection.Singular(name), tag)
		}
	} else if t == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			fillStruct(v.Index(i), inflection.Singular(name), tag)
		}
	} else {
		var fakeFn interface{}
		// if we can get fake func name from tag
		if fakeFuncName := tag.Get("fake"); fakeFuncName != "" {
			if fn, ok := fakerFuncs[strcase.ToCamel(fakeFuncName)]; ok {
				fakeFn = fn
			}
		}
		// otherwise get it from field name
		if fakeFn == nil {
			if fn, ok := fakerFuncs[strcase.ToCamel(name)]; ok {
				fakeFn = fn
			}
		}
		if fakeFn != nil {
			switch t {
			case reflect.String:
				if fn, ok := fakeFn.(func() string); ok {
					v.SetString(fn())
				}
			case reflect.Int, reflect.Int64:
				if fn, ok := fakeFn.(func() int); ok {
					v.SetInt(int64(fn()))
				}
			case reflect.Uint64:
				if fn, ok := fakeFn.(func() uint64); ok {
					v.SetUint(fn())
				}
			case reflect.Float32:
				if fn, ok := fakeFn.(func() float32); ok {
					v.SetFloat(float64(fn()))
				}
			case reflect.Float64:
				if fn, ok := fakeFn.(func() float64); ok {
					v.SetFloat(fn())
				}
			case reflect.Struct:
				if v.Type().String() == "time.Time" {
					if fn, ok := fakeFn.(func() time.Time); ok {
						v.Set(reflect.ValueOf(fn()))
					}
				}
			default:
				log.Printf("unhandled field %s as type %s\n", name, t)
			}
		} else {
			log.Println("Cannot find fake function for field ", name)
		}

	}

	return v
}
