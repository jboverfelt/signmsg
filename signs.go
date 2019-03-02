package main

import (
	"encoding/xml"
	"strings"
	"time"
)

const xmlTimeFormat = "2006-01-02T15:04:05"

type updatedTime time.Time

func (ut *updatedTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ts string
	err := d.DecodeElement(&ts, &start)
	if err != nil {
		return err
	}

	t, err := time.Parse(xmlTimeFormat, ts)
	if err != nil {
		return err
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return err
	}

	*ut = updatedTime(t.In(loc))

	return nil
}

// When we're actually displaying, just give the kitchen time
// since we're looking at it for use in the next few minutes
func (ut updatedTime) MarshalText() ([]byte, error) {
	t := time.Time(ut)
	return []byte(t.Format(time.Kitchen)), nil
}

type messageSigns struct {
	Signs []sign `xml:"messageSign"`
}

func (ms messageSigns) FindByName(names ...string) []sign {
	var signs []sign
	for _, name := range names {
		for _, mSign := range ms.Signs {
			if mSign.Name == name {
				signs = append(signs, mSign)
				break
			}
		}
	}

	return signs
}

type sign struct {
	Location  string      `xml:"location"`
	DmsID     string      `xml:"dmsid"`
	Name      string      `xml:"name"`
	Message   string      `xml:"message"`
	Updated   updatedTime `xml:"updated"`
	Beacon    bool        `xml:"beacon"`
	Latitude  string      `xml:"latitude"`
	Longitude string      `xml:"longitude"`
}

type displaySign struct {
	Location     string
	MessageLines []string
}

func toDisplaySigns(signs []sign) []displaySign {
	var displaySigns []displaySign
	for _, s := range signs {
		d := displaySign{Location: s.Location}

		noBreaks := strings.Replace(strings.ToUpper(s.Message), "<BR>", "|", -1)
		noBreaks = strings.Replace(noBreaks, "<BR/>", "|", -1)
		noBreaks = strings.Replace(noBreaks, "<BR />", "|", -1)

		d.MessageLines = strings.Split(noBreaks, "|")

		displaySigns = append(displaySigns, d)
	}

	return displaySigns
}
