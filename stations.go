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
