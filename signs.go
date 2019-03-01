package main

import (
	"encoding/xml"
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
	for _, mSign := range ms.Signs {
		for _, name := range names {
			if mSign.Name == name {
				signs = append(signs, mSign)
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
