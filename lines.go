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
		"RD": "Red",
		"YL": "Yellow",
		"GR": "Green",
		"BL": "Blue",
		"OR": "Orange",
		"SV": "Silver",
		"No": "No Passenger",
		"":   "No Passenger",
		"--": "No Passenger",
	}[l.Code]; ok {
		return name
	}
	return fmt.Sprintf("Unknown (%s)", l.Code)
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
