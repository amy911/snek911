package snek

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/amy911/amy911/onfail"
)

var (
	// The full legal text of the End User License Agreement (EULA)
	Eula string

	// The short identifier for the license as used on ChooseALicense.com
	License string = "Proprietary"
)

type Legal struct {
	Copyright Copyright `json:"copyright" xml:"copyright,attr"`
	Eula string `json:"eula" xml:"eula,attr"`
	License string `json:"license" xml:"license,attr"`
}

func NewLegal(copyright *Copyright, license, eula string) *Legal {
	return new(Legal).Init(copyright, license, eula)
}

func (legal *Legal) Init(copyright *Copyright, license, eula string) *Legal {
	legal.Copyright = *copyright
	legal.Eula = eula
	legal.License = license
	if strings.Contains(license, "\n") {
		panic("License should be the short identifier.  Use EULA for the full text.")
	}
	return legal
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
