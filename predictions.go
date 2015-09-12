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
