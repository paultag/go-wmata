package wmata

import (
	"fmt"

	"pault.ag/go/wmata/internal"
)

type Station struct {
	Code             string
	Name             string
	stationTogether1 string `json:"StationTogether1"`
	stationTogether2 string `json:"StationTogether2"`
	lineCode1        string `json:"LineCode1"`
	lineCode2        string `json:"LineCode2"`
	lineCode3        string `json:"LineCode3"`
	lineCode4        string `json:"LineCode4"`
	Lat              float64
	Lon              float64
}

func (s Station) Line() Line {
	return Line{Code: s.Code}
}

func (s Station) Lines() []Line {
	ret := []Line{}
	for _, code := range []string{
		s.lineCode1,
		s.lineCode2,
		s.lineCode3,
		s.lineCode4,
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

func GetStation(stationCode string) (Station, error) {
	target := Station{}
	err := internal.Get("Rail.svc/json/jStationInfo", map[string]string{
		"StationCode": stationCode,
	}, &target)
	return target, err
}
