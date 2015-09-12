package wmata

import (
	"fmt"
	"strings"

	"pault.ag/go/wmata/internal"
)

type Prediction struct {
	Cars            string `json:"Car"`
	Destination     string
	desitnationCode string `json:"DestinationCode"`
	DesitnationName string
	Group           string
	Line            Line
	locationCode    string `json:"LocationCode"`
	LocationName    string
	Minutes         string `json:"Min"`
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

func GetPredictions(codes ...string) (Predictions, error) {
	target := internalPredictions{}
	err := internal.Get(
		fmt.Sprintf("StationPrediction.svc/json/GetPrediction/%s", strings.Join(codes, ",")),
		map[string]string{},
		&target)
	return target.Trains, err
}

func (s Station) GetPredictions() (Predictions, error) {
	return GetPredictions(s.Code)
}
