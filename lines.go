package wmata

import (
	"fmt"
	"strings"
)

type Line struct {
	Code string
}

func (l *Line) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("Empty data coming in")
	}
	if data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("Not a string")
	}
	data = data[1 : len(data)-1]
	l.Code = strings.TrimSpace(string(data))
	return nil
}

func (l Line) Name() string {
	if name, ok := map[string]string{
		"RD": "Red Line",
		"YL": "Yellow Line",
		"GR": "Green Line",
		"BL": "Blue Line",
		"OR": "Orange Line",
		"SV": "Silver Line",
		"NO": "No Passenger",
	}[l.Code]; ok {
		return name
	}
	return fmt.Sprintf("Unknown Line: Code: %s", l.Code)
}

var RedLine = Line{Code: "RD"}
var YellowLine = Line{Code: "YL"}
var GreenLine = Line{Code: "GR"}
var BlueLine = Line{Code: "BL"}
var OrangeLine = Line{Code: "OR"}
var SilverLine = Line{Code: "SV"}

var Lines = []Line{
	RedLine, YellowLine, GreenLine,
	BlueLine, OrangeLine, SilverLine,
}
