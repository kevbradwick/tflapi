package ingest

import (
	"strings"
)

type PlatformToTrain struct {
	Line     string `xml:"trainName"`
	Platform string `xml:"platformToTrainSteps"`
}

type Entrance struct {
	Name             string            `xml:"name"`
	PlatformToTrains []PlatformToTrain `xml:"platformToTrain"`
}

type Facility struct {
	Name  string `xml:"name,attr"`
	Count string `xml:",chardata"`
}

type Station struct {
	Id          string     `xml:"id,attr"`
	Type        string     `xml:"type,attr"`
	Name        string     `xml:"name"`
	Address     string     `xml:"contactDetails>address"`
	Telephone   string     `xml:"contactDetails>phone"`
	Lines       []string   `xml:"servingLines>servingLine"`
	Zones       []string   `xml:"zones>zone"`
	Facilities  []Facility `xml:"facilities>facility"`
	Entrances   []Entrance `xml:"entrances>entrance"`
	Coordinates string     `xml:"Placemark>Point>coordinates"`
}

type Data struct {
	PublishedDate string    `xml:"Header>PublishDateTime"`
	Language      string    `xml:"Header>Language"`
	Stations      []Station `xml:"stations>station"`
}

func (s *Station) Longitude() string {
	return strings.Split(s.Coordinates, ",")[0]
}

func (s *Station) Latitude() string {
	return strings.Split(s.Coordinates, ",")[1]
}

type LatLon struct {
	Latitude, Longitude float64
}

type StationDocument struct {
	Id         string
	Name       string
	Address    string
	Telephone  string
	Lines      []string
	Zones      []string
	Facilities []Facility
	Entrances  []Entrance
	Coordinate LatLon
}