package snek

import (
	"encoding/json"
	"encoding/xml"

	"github.com/amy911/amy911/onfail"
)

var (
	// The full legal text of the End User License Agreement (EULA)
	Eula string
)

type Legal struct {
	Copyright Copyright `json:"copyright" xml:"copyright,attr"`
	Eula string `json:"eula" xml:"eula,attr"`
}

func NewLegal(copyright *Copyright, eula string) *Legal {
	return new(Legal).Init(copyright, eula)
}

func (legal *Legal) Init(copyright *Copyright, eula string) *Legal {
	legal.Copyright = *copyright
	legal.Eula = eula
}

func (legal *Legal) Json(onFail ...onfail.OnFail) string {
	out, err := json.Marshal(legal)
	if err != nil {
		onfail.Fail(err, legal, onfail.Print, onFail)
		return ""
	}
	return string(out)
}

func (legal *Legal) Xml(onFail ...onfail.OnFail) string {
	out, err := xml.Marshal(legal)
	if err != nil {
		onfail.Fail(err, legal, onfail.Print, onFail)
		return ""
	}
	return string(out)
}
