[![Build Status](https://img.shields.io/travis/icrowley/fake.svg?style=flat)](https://travis-ci.org/icrowley/fake) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/icrowley/fake) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/icrowley/fake/master/LICENSE)

Updates: Oct 11, 2018 

1. added Chinese language data and functions.
2. FillStruct automatically fill struct field with fake data by tag or fieldnames. See examples section for details.

Fake
====

Fake is a fake data generator for Go (Golang), heavily inspired by the forgery and ffaker Ruby gems.

## About

Most data and methods are ported from forgery/ffaker Ruby gems.
For the list of available methods please look at https://godoc.org/github.com/icrowley/fake.
Currently english and russian languages are available.

Fake embeds samples data files unless you call `UseExternalData(true)` in order to be able to work without external files dependencies when compiled, so, if you add new data files or make changes to existing ones don't forget to regenerate data.go file using `github.com/mjibson/esc` tool and `esc -o data.go -pkg fake data` command (or you can just use `go generate` command if you are using Go 1.4 or later).

## Install

```shell
go get github.com/icrowley/fake
```

## Import

```go
import (
  "github.com/icrowley/fake"
)
```

## Documentation

Documentation can be found at godoc:

https://godoc.org/github.com/icrowley/fake

## Test
To run the project tests:

```shell
cd test
go test
```

## Examples

```go
name := fake.FirstName()
fullname := fake.FullName()
product := fake.Product()
```

Changing language:

```go
err := fake.SetLang("ru")
if err != nil {
  panic(err)
}
password := fake.SimplePassword()
```

Using english fallback:

```go
err := fake.SetLang("ru")
if err != nil {
  panic(err)
}
fake.EnFallback(true)
password := fake.Paragraph()
```

Using external data:

```go
fake.UseExternalData(true)
password := fake.Paragraph()
```

Using FillStruct:

 FillStruct fills struct field with faked data.  At first FillStruct will try to get the field tag with tag key "fake" for which fake function to call. If "fake" tag is not set, it will look at the field name if matches a fake function name. Otherwise a error is logged out.

```go
type A struct {
    UserName string
    Profile  struct {
        EmailAddresses [2]string
        Place          string `fake:"StreetAddress"`
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
    Sex            string
    Languages      []string
    Brand          string
    ProductName    string
    Product        string
    URL            string
    Description    string
    Price          uint64
    Quantity       uint64
    PriceF         float64
    CreatedAt      time.Time
}
a := &A{}
a.Languages = make([]string, 2)

FillStruct(&a)
fmt.Printf("%+v\n", *a)
```

### Author

Dmitry Afanasyev,
http://twitter.com/i_crowley
dimarzio1986@gmail.com


### Maintainers

Dmitry Moskowski
https://github.com/corpix
