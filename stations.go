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

	"pault.ag/go/wmata/internal"
)

type Station struct {
	Code             string
	Name             string
	StationTogether1 string
	StationTogether2 string
	LineCode1        string
	LineCode2        string
	LineCode3        string
	LineCode4        string
	Lat              float64
	Lon              float64
}

func (s Station) Line() Line {
	return Line{Code: s.Code}
}

func (s Station) Lines() []Line {
	ret := []Line{}
	for _, code := range []string{
		s.LineCode1,
		s.LineCode2,
		s.LineCode3,
		s.LineCode4,
	} {
		ret = append(ret, Line{Code: code})
	}
	return ret
}

func (s Station) String() string {
	return fmt.Sprintf("%s (%s)", s.Name, s.Code)
}

type Stations struct {
	Stations []Station
}

func (l Line) GetStations() ([]Station, error) {
	target := Stations{}
	err := internal.Get("Rail.svc/json/jStations", map[string]string{
		"LineCode": l.Code,
	}, &target)
	return target.Stations, err
}

func GetStations() ([]Station, error) {
	target := Stations{}
	err := internal.Get("Rail.svc/json/jStations", map[string]string{}, &target)
	return target.Stations, err
}

func GetStation(stationCode string) (*Station, error) {
	target := Station{}
	err := internal.Get("Rail.svc/json/jStationInfo", map[string]string{
		"StationCode": stationCode,
	}, &target)
	return &target, err
}

// vim: foldmethod=marker
