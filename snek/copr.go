package snek

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/amy911/amy911/onfail"
)

var (
	// The first year that the copyright was held
	CopyrightFirstYear int

	// The holder of the copyright
	CopyrightHolder string
)

type Copyright struct {
	Holder string `json:"holder" xml:"holder,attr"`
	Pretty string `json:"pretty" xml:"pretty,attr"`
	Robots string `json:"robots" xml:"robots,attr"`
	From   int    `json:"from" xml:"from,attr"`
	To     int    `json:"to" xml:"to,attr"`
}

func NewCopyright(firstYear int, holder string) *Copyright {
	return new(Copyright).Init(firstYear, holder)
}

func (copyright *Copyright) Init(firstYear int, holder string) *Copyright {
	thisYear := time.Now().Year()
	years := strconv.Itoa(firstYear) + "-" + strconv.Itoa(thisYear)
	copyright.Holder = holder
	copyright.Pretty = "Copyright \u00a9 " + years + " " + holder
	copyright.Robots = "Copyright (c) " + years + "\t" + holder
	copyright.From = firstYear
	copyright.To = thisYear
	return new(Copyright).Init(firstYear, holder)
}

func (copyright *Copyright) Json(onFail ...onfail.OnFail) string {
	out, err := json.Marshal(copyright)
	if err != nil {
		onfail.Fail(err, copyright, onfail.Print, onFail)
		return ""
	}
	return string(out)
}

func (copyright *Copyright) Xml(onFail ...onfail.OnFail) string {
	out, err := xml.Marshal(copyright)
	if err != nil {
		onfail.Fail(err, copyright, onfail.Print, onFail)
		return ""
	}
	return string(out)
}
