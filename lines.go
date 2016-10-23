// {{{ Copyright (c) Paul R. Tagliamonte <paultag@dc.cant.vote>, 2015
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

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

var LineMap = map[string]Line{
	"RD": RedLine,
	"YL": YellowLine,
	"GR": GreenLine,
	"BL": BlueLine,
	"OR": OrangeLine,
	"SV": SilverLine,
}

// vim: foldmethod=marker
