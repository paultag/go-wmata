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

	"pault.ag/go/wmata/internal"
)

type Prediction struct {
	Cars        string `json:"Car"`
	Destination string
	Group       string
	Line        Line
	Minutes     string `json:"Min"`

	DesitnationName string
	DesitnationCode string

	LocationName string
	LocationCode string
}

func (p Prediction) GetDestination() (*Station, error) {
	if p.DesitnationCode == "" {
		return nil, fmt.Errorf("No known desitnation")
	}
	return GetStation(p.DesitnationCode)
}

func (p Prediction) GetLocation() (*Station, error) {
	fmt.Printf("%s\n", p.LocationCode)
	if p.LocationCode == "" {
		return nil, fmt.Errorf("No known location")
	}
	return GetStation(p.LocationCode)
}

func (p Prediction) String() string {
	return fmt.Sprintf(
		"%s line %s Car Train to %s, arriving at %s in %s minutes",
		p.Line.Name(),
		p.Cars,
		p.Destination,
		p.LocationName,
		p.Minutes,
	)
}

type internalPredictions struct {
	Trains Predictions
}

type Predictions []Prediction

func (predictions Predictions) Group(group string) Predictions {
	ret := Predictions{}
	for _, prediction := range predictions {
		if prediction.Group == group {
			ret = append(ret, prediction)
		}
	}
	return ret
}

func GetPredictionsByCodes(codes ...string) (Predictions, error) {
	target := internalPredictions{}
	err := internal.Get(
		fmt.Sprintf("StationPrediction.svc/json/GetPrediction/%s", strings.Join(codes, ",")),
		map[string]string{},
		&target)
	return target.Trains, err
}

func (s Station) GetPredictions() (Predictions, error) {
	return GetPredictionsByCodes(s.Code)
}

// vim: foldmethod=marker
