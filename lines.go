package wmata

import (
	"fmt"
)

type Line struct {
	Code string
}

func (l Line) Name() string {
	if name, present := map[string]string{
		"RD": "Red Line",
		"YL": "Yellow Line",
		"GR": "Green Line",
		"BL": "Blue Line",
		"OR": "Orange Line",
		"SV": "Silver Line",
		"NO": "No Passenger",
	}[l.Code]; present {
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
