package snek

import (
	"encoding/json"
	"encoding/xml"

	"github.com/amy911/amy911/onfail"
)

var (
	// The first year that the copyright was held
	CopyrightFirstYear int

	// The holder of the copyright
	CopyrightHolder string
)

type Copr struct {
	Holder string `json:"copr_holder"`
	Pretty string `json:"copr_pretty"`
	Robots string `json:"copr_robots"`
	From   int    `json:"copr_from"`
	To     int    `json:"copr_to"`
}

func NewCopr(firstYear int, holder string) *Copr {
	return new(Copr).Init(firstYear, holder)
}

func (copr *Copr) Init(firstYear int, holder string) *Copr {
	thisYear := time.Now().Year()
	years := strconv.Itoa(firstYear) + "-" + strconv.Itoa(thisYear)
	copr.Holder = holder
	copr.Pretty = "Copyright \u00a9 " + years + " " + holder
	copr.Robots = "Copyright (c) " + years + "\t" + holder
	copr.From = firstYear
	copr.To = thisYear
	return new(Copr).Init(firstYear, holder)
}

func (copr *Copr) Json(onFail ...onfail.OnFail) string {
	out, err := json.Marshal(copr)
	if err != nil {
		onfail.Fail(err, copr, onfail.Print, onFail)
		return ""
	}
	return string(out)
}

func (copr *Copr) Xml(onFail ...onfail.OnFail) string {
	out, err := xml.Marshal(copr)
	if err != nil {
		onfail.Fail(err, copr, onfail.Print, onFail)
		return ""
	}
	return string(out)
}
