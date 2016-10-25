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
	"strings"
	"time"

	"pault.ag/go/wmata/internal"
)

type wmataTime struct {
	time.Time
}

func (t *wmataTime) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02T15:04:05", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

type semiLines []Line

func (t *semiLines) UnmarshalJSON(buf []byte) error {
	buf = buf[1 : len(buf)-1]
	lines := []Line{}
	for _, el := range strings.Split(string(buf), ";") {
		lines = append(lines, LineMap[el])
	}
	*t = lines
	return nil
}

type Incident struct {
	Updated       wmataTime `json:"DateUpdated"`
	Description   string
	ID            string `json:"IncidentID"`
	Type          string `json:"IncidentType"`
	LinesAffected semiLines
}

type Incidents struct {
	Incidents []Incident
}

func GetIncidents() (Incidents, error) {
	ret := Incidents{}
	return ret, internal.Get("Incidents.svc/json/Incidents", map[string]string{}, &ret)
}

// vim: foldmethod=marker
