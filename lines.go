package wmata

import (
	"fmt"
	"pault.ag/go/wmata/internal"
)

type line string

var RedLine line = "RD"
var YellowLine line = "YL"
var GreenLine line = "GR"
var BlueLine line = "BL"
var OrangeLine line = "OR"
var SilverLine line = "SV"

type Station struct {
	Code             string
	Lat              float64
	Lon              float64
	LineCode1        string
	LineCode2        string
	LineCode3        string
	LineCode4        string
	Name             string
	StationTogether1 string
	StationTogether2 string
}

type Prediction struct {
	Car             string
	Destination     string
	DestinationCode string
	DestinationName string
	Group           string
	Line            line
	LocationCode    string
	LocationName    string
	Min             string
}

func (p Prediction) String() string {
	return fmt.Sprintf(
		"%s car train to %s arriving at %s in %s (Group %s)",
		p.Car,
		p.DestinationCode,
		p.LocationCode,
		p.Min,
		p.Group,
	)
}

type predictions struct {
	Trains []Prediction
}

func (s Station) Predictions() ([]Prediction, error) {
	target := predictions{}
	err := internal.Get(
		fmt.Sprintf("StationPrediction.svc/json/GetPrediction/%s", s.Code),
		map[string]string{},
		&target,
	)
	return target.Trains, err
}

type stations struct {
	Stations []Station
}

func Stations(whatLine line) (map[string]Station, error) {
	target := stations{}
	err := internal.Get("Rail.svc/json/jStations", map[string]string{
		"LineCode": string(whatLine),
	}, &target)
	ret := map[string]Station{}
	if err != nil {
		return ret, err
	}
	for _, station := range target.Stations {
		ret[station.Code] = station
	}
	return ret, err
}
